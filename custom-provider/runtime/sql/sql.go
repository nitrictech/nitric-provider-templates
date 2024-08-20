package sql

import (
	"context"

	sqlpb "github.com/nitrictech/nitric/core/pkg/proto/sql/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SqlServer struct{}

var _ sqlpb.SqlServer = (*SqlServer)(nil)

func (s *SqlServer) ConnectionString(ctx context.Context, req *sqlpb.SqlConnectionStringRequest) (*sqlpb.SqlConnectionStringResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func New() (*SqlServer, error) {
	return &SqlServer{}, nil
}
