package deploy

import (
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func (a *AwsExtensionProvider) Service(ctx *pulumi.Context, parent pulumi.Resource, name string, config *pulumix.NitricPulumiServiceConfig, runtime provider.RuntimeProvider) error {
	// Append Digital Ocean environment variables
	config.SetEnv("DIGITALOCEAN_REGION", pulumi.String(a.config.Spaces.Region))
	config.SetEnv("SPACES_KEY", pulumi.String(a.config.Spaces.Key))
	config.SetEnv("SPACES_SECRET", pulumi.String(a.config.Spaces.Secret))

	return a.NitricAwsPulumiProvider.Service(ctx, parent, name, config, runtime)
}
