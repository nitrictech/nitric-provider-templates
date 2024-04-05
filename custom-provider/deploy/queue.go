package deploy

import (
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Queue struct {
	pulumi.ResourceState
	Name string
}

func (n *NitricCustomPulumiProvider) Queue(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Queue) error {
	return nil
}
