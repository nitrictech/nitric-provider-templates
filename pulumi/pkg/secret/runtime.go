package secret

import (
	"context"

	secretpb "github.com/nitrictech/nitric/core/pkg/proto/secret/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SecretServer struct{}

var _ secretpb.SecretServiceServer = &SecretServer{}

// Updates a secret, creating a new one if it doesn't already exist
func (srv *SecretServer) Put(ctx context.Context, req *secretpb.SecretPutRequest) (*secretpb.SecretPutResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Gets a secret from a Secret Store
func (srv *SecretServer) Access(ctx context.Context, req *secretpb.SecretAccessRequest) (*secretpb.SecretAccessResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func NewServer() *SecretServer {
	return &SecretServer{}
}
