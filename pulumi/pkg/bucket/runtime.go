package bucket

import (
	"context"

	storagepb "github.com/nitrictech/nitric/core/pkg/proto/storage/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BucketServer struct{}

var _ storagepb.StorageServiceServer = &BucketServer{}

// Retrieve an item from a bucket
func (*BucketServer) Read(context.Context, *storagepb.StorageReadRequest) (*storagepb.StorageReadResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Store an item to a bucket
func (*BucketServer) Write(context.Context, *storagepb.StorageWriteRequest) (*storagepb.StorageWriteResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Delete an item from a bucket
func (*BucketServer) Delete(context.Context, *storagepb.StorageDeleteRequest) (*storagepb.StorageDeleteResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Generate a pre-signed URL for direct operations on an item
func (*BucketServer) PreSignUrl(context.Context, *storagepb.StoragePreSignUrlRequest) (*storagepb.StoragePreSignUrlResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// List files currently in the bucket
func (*BucketServer) ListFiles(context.Context, *storagepb.StorageListFilesRequest) (*storagepb.StorageListFilesResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Determine is an object exists in a bucket
func (*BucketServer) Exists(context.Context, *storagepb.StorageExistsRequest) (*storagepb.StorageExistsResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func NewServer() *BucketServer {
	return &BucketServer{}
}
