package deploy

import (
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This represents a unique unit of execution at the moment this is a container but could also be many things e.g. WASM, Binary, source zip etc.
type Service struct {
	pulumi.ResourceState
	Name string
}

func (a *NitricCustomPulumiProvider) Service(ctx *pulumi.Context, parent pulumi.Resource, name string, config *pulumix.NitricPulumiServiceConfig, provider provider.RuntimeProvider) error {
	return nil
}
