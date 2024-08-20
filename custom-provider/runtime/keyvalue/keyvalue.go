package keyvalue

import (
	"context"

	kvstorepb "github.com/nitrictech/nitric/core/pkg/proto/kvstore/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KvStoreServer struct{}

var _ kvstorepb.KvStoreServer = &KvStoreServer{}

// Updates a secret, creating a new one if it doesn't already exist
// Get an existing document
func (*KvStoreServer) GetValue(context.Context, *kvstorepb.KvStoreGetValueRequest) (*kvstorepb.KvStoreGetValueResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Create a new or overwrite an existing document
func (*KvStoreServer) SetValue(context.Context, *kvstorepb.KvStoreSetValueRequest) (*kvstorepb.KvStoreSetValueResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Delete an existing document
func (*KvStoreServer) DeleteKey(context.Context, *kvstorepb.KvStoreDeleteKeyRequest) (*kvstorepb.KvStoreDeleteKeyResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Iterate over all keys in a store
func (*KvStoreServer) ScanKeys(*kvstorepb.KvStoreScanKeysRequest, kvstorepb.KvStore_ScanKeysServer) error {
	return status.New(codes.Unimplemented, "Unimplemented").Err()
}

func New() (*KvStoreServer, error) {
	return &KvStoreServer{}, nil
}
