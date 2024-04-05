package api

import (
	"context"

	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/resource"
	apipb "github.com/nitrictech/nitric/core/pkg/proto/apis/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiServer struct{}

var _ apipb.ApiServer = &ApiServer{}

// ApiDetails implements apispb.ApiServer.
func (*ApiServer) ApiDetails(context.Context, *apipb.ApiDetailsRequest) (*apipb.ApiDetailsResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Serve implements apispb.ApiServer.
func (*ApiServer) Serve(apipb.Api_ServeServer) error {
	return status.New(codes.Unimplemented, "Unimplemented").Err()
}

func New(provider *resource.ResourceServer) (*ApiServer, error) {
	return &ApiServer{}, nil
}
