package deploy

import (
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	resourcespb "github.com/nitrictech/nitric/core/pkg/proto/resources/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/samber/lo"
)

func (a *AwsExtensionProvider) Policy(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Policy) error {
	filteredConfig := deploymentspb.Policy{
		Principals: config.Principals,
	}

	filteredConfig.Resources = lo.Filter(config.Resources, func(res *deploymentspb.Resource, idx int) bool {
		return res.Id.Type != resourcespb.ResourceType_Bucket
	})

	filteredConfig.Actions = lo.Filter(config.Actions, func(res resourcespb.Action, idx int) bool {
		return !lo.Contains([]resourcespb.Action{
			resourcespb.Action_BucketFileDelete,
			resourcespb.Action_BucketFileGet,
			resourcespb.Action_BucketFileList,
			resourcespb.Action_BucketFilePut,
		}, res)
	})

	if len(filteredConfig.Actions) == 0 {
		return nil
	}

	return a.NitricAwsPulumiProvider.Policy(ctx, parent, name, &filteredConfig)
}
