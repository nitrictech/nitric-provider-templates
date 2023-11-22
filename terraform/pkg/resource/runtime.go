package resource

import (
	"context"

	resourcepb "github.com/nitrictech/nitric/core/pkg/proto/resource/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResourceServer struct{}

var _ resourcepb.ResourceServiceServer = &ResourceServer{}

// Declare a resource for the nitric application
// At Deploy time this will create resources as part of the nitric stacks dependency graph
// At runtime
func (*ResourceServer) Declare(context.Context, *resourcepb.ResourceDeclareRequest) (*resourcepb.ResourceDeclareResponse, error) {
	// FIXME: This is normally a no-op for runtime but could be used to eagerly resolve resource references on startup
	return &resourcepb.ResourceDeclareResponse{}, nil
}

// Retrieve details about a resource at runtime
func (*ResourceServer) Details(context.Context, *resourcepb.ResourceDetailsRequest) (*resourcepb.ResourceDetailsResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}
