package deploy

import (
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schedule struct {
	pulumi.ResourceState
	Name string
}

func (a *NitricCustomPulumiProvider) Schedule(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Schedule) error {
	return nil
}
