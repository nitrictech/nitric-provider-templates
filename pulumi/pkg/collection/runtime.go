package collection

import (
	"context"

	documentpb "github.com/nitrictech/nitric/core/pkg/proto/document/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CollectionServer struct{}

var _ documentpb.DocumentServiceServer = &CollectionServer{}

// Updates a secret, creating a new one if it doesn't already exist
// Get an existing document
func (*CollectionServer) Get(context.Context, *documentpb.DocumentGetRequest) (*documentpb.DocumentGetResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Create a new or overwrite an existing document
func (*CollectionServer) Set(context.Context, *documentpb.DocumentSetRequest) (*documentpb.DocumentSetResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Delete an existing document
func (*CollectionServer) Delete(context.Context, *documentpb.DocumentDeleteRequest) (*documentpb.DocumentDeleteResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Query the document collection (supports pagination)
func (*CollectionServer) Query(context.Context, *documentpb.DocumentQueryRequest) (*documentpb.DocumentQueryResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Query the document collection (supports streaming)
func (*CollectionServer) QueryStream(*documentpb.DocumentQueryStreamRequest, documentpb.DocumentService_QueryStreamServer) error {
	return status.New(codes.Unimplemented, "Unimplemented").Err()
}

func NewServer() *CollectionServer {
	return &CollectionServer{}
}
