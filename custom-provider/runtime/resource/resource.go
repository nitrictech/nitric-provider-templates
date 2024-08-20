package resource

import (
	"context"

	resourcepb "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResourceServer struct{}

var _ resourcepb.ResourcesServer = &ResourceServer{}

// Declare a resource for the nitric application
// At Deploy time this will create resources as part of the nitric stacks dependency graph
// At runtime
func (*ResourceServer) Declare(context.Context, *resourcepb.ResourceDeclareRequest) (*resourcepb.ResourceDeclareResponse, error) {
	// FIXME: This is normally a no-op for runtime but could be used to eagerly resolve resource references on startup
	return &resourcepb.ResourceDeclareResponse{}, nil
}

// Retrieve details about a resource at runtime
func (*ResourceServer) Details(context.Context, *resourcepb.ResourceDeclareRequest) (*resourcepb.ResourceDeclareResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func New() (*ResourceServer, error) {
	return &ResourceServer{}, nil
}
