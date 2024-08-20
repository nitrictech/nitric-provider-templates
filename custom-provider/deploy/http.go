package deploy

import (
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *NitricCustomPulumiProvider) Http(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Http) error {
	// TODO: Implement HTTP deployment for the custom provider
	return status.Error(codes.Unimplemented, "HTTP Proxy deployment is not implemented for the custom provider")
}
