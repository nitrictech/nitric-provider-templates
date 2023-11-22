package main

import (
	"fmt"
	"log"
	"net"

	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/nitrictech/nitric/core/pkg/utils"
	"google.golang.org/grpc"
)

type DeployServer struct {
	deploypb.UnimplementedDeployServiceServer
}

// Start the deployment server
func main() {
	// TODO: Most of this block is already standardised as part of our existing providers
	port := utils.GetEnv("PORT", "50051")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("error listening on port %s %v", port, err)
	}

	srv := grpc.NewServer()

	deploypb.RegisterDeployServiceServer(srv, &DeployServer{})

	fmt.Printf("Deployment server started on %s\n", lis.Addr().String())
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("error serving requests %v", err)
	}
}
