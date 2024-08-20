package storage

import (
	"context"

	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StorageServer struct{}

var _ storagepb.StorageServer = &StorageServer{}

// Retrieve an item from a bucket
func (*StorageServer) Read(context.Context, *storagepb.StorageReadRequest) (*storagepb.StorageReadResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Store an item to a bucket
func (*StorageServer) Write(context.Context, *storagepb.StorageWriteRequest) (*storagepb.StorageWriteResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Delete an item from a bucket
func (*StorageServer) Delete(context.Context, *storagepb.StorageDeleteRequest) (*storagepb.StorageDeleteResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Generate a pre-signed URL for direct operations on an item
func (*StorageServer) PreSignUrl(context.Context, *storagepb.StoragePreSignUrlRequest) (*storagepb.StoragePreSignUrlResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// List files currently in the bucket
func (*StorageServer) ListBlobs(context.Context, *storagepb.StorageListBlobsRequest) (*storagepb.StorageListBlobsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Determine is an object exists in a bucket
func (*StorageServer) Exists(context.Context, *storagepb.StorageExistsRequest) (*storagepb.StorageExistsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func New() (*StorageServer, error) {
	return &StorageServer{}, nil
}
