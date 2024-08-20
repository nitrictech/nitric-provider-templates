package deploy

import (
	"github.com/nitrictech/nitric/cloud/aws/deploy"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtendedAwsProvider struct {
	deploy.NitricAwsPulumiProvider

	// Add additional fields here, such as custom configuration or resource references
}

// Implement replacements for any resource creation steps you like, such as Topics, Buckets, Queues, etc.
func (a *ExtendedAwsProvider) Topic(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Topic) error {
	// Replace this line with your custom implementation
	// 	Explore the existing provider implementation on GitHub for examples.
	return a.NitricAwsPulumiProvider.Topic(ctx, parent, name, config)
}

func NewExtendedAwsProvider() *ExtendedAwsProvider {
	baseProvider := deploy.NewNitricAwsProvider()

	return &ExtendedAwsProvider{
		NitricAwsPulumiProvider: *baseProvider,
	}
}
