package api

import (
	"context"

	apipb "github.com/nitrictech/nitric/core/pkg/proto/apis/v1"
	"github.com/nitrictech/nitric/core/pkg/workers/apis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiServer struct {
	*apis.RouteWorkerManager
}

var _ apipb.ApiServer = &ApiServer{}

func (*ApiServer) ApiDetails(context.Context, *apipb.ApiDetailsRequest) (*apipb.ApiDetailsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func New() (*ApiServer, error) {
	return &ApiServer{}, nil
}
