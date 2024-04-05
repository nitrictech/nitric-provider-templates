// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"google.golang.org/grpc/codes"

	"github.com/nitrictech/nitric/cloud/aws/ifaces/s3iface"
	"github.com/nitrictech/nitric/cloud/aws/runtime/env"
	"github.com/nitrictech/nitric/cloud/aws/runtime/resource"
	grpc_errors "github.com/nitrictech/nitric/core/pkg/grpc/errors"
	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
)

// SpacesStorageService - an AWS S3 implementation of the Nitric Storage Service
type SpacesStorageService struct {
	s3Client      s3iface.S3API
	preSignClient s3iface.PreSignAPI
	provider      resource.AwsResourceProvider
}

var _ storagepb.StorageServer = (*SpacesStorageService)(nil)

func isS3AccessDeniedErr(err error) bool {
	var opErr *smithy.OperationError
	if errors.As(err, &opErr) {
		return opErr.Service() == "S3" && strings.Contains(opErr.Unwrap().Error(), "AccessDenied")
	}

	return false
}

// Read and return the contents of a file in a bucket
func (s *SpacesStorageService) Read(ctx context.Context, req *storagepb.StorageReadRequest) (*storagepb.StorageReadResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("S3StorageService.Read")

	bucketName := &req.BucketName

	resp, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: bucketName,
		Key:    aws.String(req.Key),
	})
	if err != nil {
		if isS3AccessDeniedErr(err) {
			return nil, newErr(
				codes.PermissionDenied,
				"unable to read file, this may be due to a missing permissions request in your code.",
				err,
			)
		}

		return nil, newErr(
			codes.Unknown,
			"error reading file",
			err,
		)
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &storagepb.StorageReadResponse{
		Body: bodyBytes,
	}, nil
}

// Write contents to a file in a bucket
func (s *SpacesStorageService) Write(ctx context.Context, req *storagepb.StorageWriteRequest) (*storagepb.StorageWriteResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("S3StorageService.Write")

	bucketName := &req.BucketName

	contentType := http.DetectContentType(req.Body)

	if _, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      bucketName,
		Body:        bytes.NewReader(req.Body),
		ContentType: &contentType,
		Key:         aws.String(req.Key),
	}); err != nil {
		if isS3AccessDeniedErr(err) {
			return nil, newErr(
				codes.PermissionDenied,
				"unable to write file, this may be due to a missing permissions request in your code.",
				err,
			)
		}

		return nil, newErr(
			codes.Unknown,
			"error writing file",
			err,
		)
	}

	return &storagepb.StorageWriteResponse{}, nil
}

// Delete a file from a bucket
func (s *SpacesStorageService) Delete(ctx context.Context, req *storagepb.StorageDeleteRequest) (*storagepb.StorageDeleteResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("S3StorageService.Delete")

	bucketName := &req.BucketName

	if _, err := s.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: bucketName,
		Key:    aws.String(req.Key),
	}); err != nil {
		if isS3AccessDeniedErr(err) {
			return nil, newErr(
				codes.PermissionDenied,
				"unable to delete file, this may be due to a missing permissions request in your code.",
				err,
			)
		}

		return nil, newErr(
			codes.Unknown,
			"error deleting file",
			err,
		)
	}

	return &storagepb.StorageDeleteResponse{}, nil
}

// PreSignUrl generates a signed URL which can be used to perform direct operations on a file
// useful for large file uploads/downloads so they can bypass application code and work directly with S3
func (s *SpacesStorageService) PreSignUrl(ctx context.Context, req *storagepb.StoragePreSignUrlRequest) (*storagepb.StoragePreSignUrlResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("S3StorageService.PreSignUrl")

	bucketName := &req.BucketName

	switch req.Operation {
	case storagepb.StoragePreSignUrlRequest_READ:
		response, err := s.preSignClient.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket: bucketName,
			Key:    aws.String(req.Key),
		}, s3.WithPresignExpires(req.Expiry.AsDuration()))
		if err != nil {
			return nil, newErr(
				codes.Internal,
				"failed to generate signed READ URL",
				err,
			)
		}
		return &storagepb.StoragePreSignUrlResponse{
			Url: response.URL,
		}, err
	case storagepb.StoragePreSignUrlRequest_WRITE:
		req, err := s.preSignClient.PresignPutObject(ctx, &s3.PutObjectInput{
			Bucket: bucketName,
			Key:    aws.String(req.Key),
		}, s3.WithPresignExpires(req.Expiry.AsDuration()))
		if err != nil {
			return nil, newErr(
				codes.Internal,
				"failed to generate signed WRITE URL",
				err,
			)
		}
		return &storagepb.StoragePreSignUrlResponse{
			Url: req.URL,
		}, err
	default:
		return nil, newErr(codes.Unimplemented, "requested operation not supported for pre-signed AWS S3 URLs", nil)
	}
}

// ListFiles lists all files in a bucket
func (s *SpacesStorageService) ListBlobs(ctx context.Context, req *storagepb.StorageListBlobsRequest) (*storagepb.StorageListBlobsResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("S3StorageService.ListFiles")

	var prefix *string = nil
	if req.Prefix != "" {
		// Only apply if prefix isn't default
		prefix = &req.Prefix
	}

	bucketName := &req.BucketName

	objects, err := s.s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: bucketName,
		Prefix: prefix,
	})
	if err != nil {
		if isS3AccessDeniedErr(err) {
			return nil, newErr(
				codes.PermissionDenied,
				"unable to list files, this may be due to a missing permissions request in your code.",
				err,
			)
		}

		return nil, newErr(
			codes.Unknown,
			"error listing files",
			err,
		)
	}

	files := make([]*storagepb.Blob, 0, len(objects.Contents))
	for _, o := range objects.Contents {
		files = append(files, &storagepb.Blob{
			Key: *o.Key,
		})
	}

	return &storagepb.StorageListBlobsResponse{
		Blobs: files,
	}, nil
}

func (s *SpacesStorageService) Exists(ctx context.Context, req *storagepb.StorageExistsRequest) (*storagepb.StorageExistsResponse, error) {
	newErr := grpc_errors.ErrorsWithScope("S3StorageService.Exists")

	bucketName := &req.BucketName

	_, err := s.s3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: bucketName,
		Key:    aws.String(req.Key),
	})
	if err != nil {
		if isS3AccessDeniedErr(err) {
			return nil, newErr(
				codes.PermissionDenied,
				"unable to check if file exists, this may be due to a missing permissions request in your code.",
				err,
			)
		}

		return &storagepb.StorageExistsResponse{
			Exists: false,
		}, nil
	}

	return &storagepb.StorageExistsResponse{
		Exists: true,
	}, nil
}

// New creates a new default S3 storage plugin
func New(provider resource.AwsResourceProvider) (*SpacesStorageService, error) {
	awsRegion := env.AWS_REGION.String()
	doRegion := os.Getenv("DIGITALOCEAN_REGION")
	accessKey := os.Getenv("SPACES_KEY")
	secretKey := os.Getenv("SPACES_SECRET")
	spacesEndpoint := fmt.Sprintf("https://%s.digitaloceanspaces.com", doRegion)

	s3Client := s3.New(s3.Options{
		Region:       awsRegion,
		Credentials:  aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		BaseEndpoint: &spacesEndpoint,
	})
	return &SpacesStorageService{
		s3Client:      s3Client,
		preSignClient: s3.NewPresignClient(s3Client),
		provider:      provider,
	}, nil
}
