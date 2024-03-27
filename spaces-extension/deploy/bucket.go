package deploy

import (
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replace the key value store configuration
func (a *AwsExtensionProvider) Bucket(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Bucket) error {
	bucket, err := digitalocean.NewSpacesBucket(ctx, name, &digitalocean.SpacesBucketArgs{
		Name:   pulumi.String(name),
		Region: pulumi.String(a.config.Spaces.Region),
		Acl:    pulumi.String("private"),
	})
	if err != nil {
		return err
	}

	a.Buckets[name] = bucket

	return nil
}
