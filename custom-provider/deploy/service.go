package deploy

import (
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *NitricCustomPulumiProvider) Service(ctx *pulumi.Context, parent pulumi.Resource, name string, config *pulumix.NitricPulumiServiceConfig, runtime provider.RuntimeProvider) error {
	// TODO: Implement Service deployment for the custom provider
	return status.Error(codes.Unimplemented, "Service deployment is not implemented for the custom provider")
}
