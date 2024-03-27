package deploy

import (
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Http struct {
	pulumi.ResourceState
	Name string
}

func (n *NitricCustomPulumiProvider) Http(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Http) error {
	return nil
}
