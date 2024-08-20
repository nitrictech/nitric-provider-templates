package deploy

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NitricCustomPulumiProvider struct {
	*deploy.CommonStackDetails

	StackId string

	config *CustomConfig

	provider.NitricDefaultOrder
}

var _ provider.NitricPulumiProvider = (*NitricCustomPulumiProvider)(nil)

func (a *NitricCustomPulumiProvider) Config() (auto.ConfigMap, error) {
	return auto.ConfigMap{
		"docker:version": auto.ConfigValue{Value: deploy.PulumiDockerVersion},
	}, nil
}

func (a *NitricCustomPulumiProvider) Init(attributes map[string]interface{}) error {
	var err error

	a.CommonStackDetails, err = deploy.CommonStackDetailsFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	a.config, err = ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "bad stack configuration: %s", err)
	}

	return nil
}

func (a *NitricCustomPulumiProvider) Pre(ctx *pulumi.Context, resources []*pulumix.NitricPulumiResource[any]) error {
	// Implement pre-deployment logic here, typically this is a good place to generate a random stack id

	return nil
}

func (a *NitricCustomPulumiProvider) Post(ctx *pulumi.Context) error {
	return nil
}

func (a *NitricCustomPulumiProvider) Result(ctx *pulumi.Context) (pulumi.StringOutput, error) {
	outputs := []interface{}{}

	output, ok := pulumi.All(outputs...).ApplyT(func(details []interface{}) string {
		stringyOutputs := make([]string, len(details))
		for i, d := range details {
			stringyOutputs[i] = d.(string)
		}

		return strings.Join(stringyOutputs, "\n")
	}).(pulumi.StringOutput)

	if !ok {
		return pulumi.StringOutput{}, fmt.Errorf("failed to generate pulumi output")
	}

	return output, nil
}

func NewNitricCustomPulumiProvider() *NitricCustomPulumiProvider {
	return &NitricCustomPulumiProvider{}
}
