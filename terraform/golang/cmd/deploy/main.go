package main

import (
	"log"
	"net"

	"github.com/nitrictech/nitric-provider-templates/terraform/golang/deploy"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"google.golang.org/grpc"
)

func main() {
	// start the deployment grpc server
	server := deploy.NewDeploymentServer()

	srv := grpc.NewServer()

	deploymentspb.RegisterDeploymentServer(srv, server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	srv.Serve(lis)
}
