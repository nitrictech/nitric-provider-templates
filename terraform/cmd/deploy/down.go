package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
)

func (*DeployServer) Down(req *deploypb.DeployDownRequest, stream deploypb.DeployService_DownServer) error {
	return status.Errorf(codes.Unimplemented, "Down not implemented")
}
