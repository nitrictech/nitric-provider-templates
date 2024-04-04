package deploy

import (
	"fmt"

	"github.com/nitrictech/nitric/cloud/aws/deploy"
	common "github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	pulumiAwsVersion = "6.6.0"
	pulumiDOVersion  = "4.27.0"
)

type AwsExtensionProvider struct {
	deploy.NitricAwsPulumiProvider

	config *ExtensionConfig

	Buckets map[string]*digitalocean.SpacesBucket
}

func NewAwsExtensionProvider() *AwsExtensionProvider {
	awsProvider := deploy.NewNitricAwsProvider()

	return &AwsExtensionProvider{
		NitricAwsPulumiProvider: *awsProvider,
		Buckets:                 make(map[string]*digitalocean.SpacesBucket),
	}
}

func (a *AwsExtensionProvider) Config() (auto.ConfigMap, error) {
	return auto.ConfigMap{
		"aws:region":                     auto.ConfigValue{Value: a.Region},
		"aws:version":                    auto.ConfigValue{Value: pulumiAwsVersion},
		"digitalocean:token":             auto.ConfigValue{Value: a.config.Token},
		"digitalocean:spaces_access_id":  auto.ConfigValue{Value: a.config.Spaces.Key},
		"digitalocean:spaces_secret_key": auto.ConfigValue{Value: a.config.Spaces.Secret},
		"digitalocean:version":           auto.ConfigValue{Value: pulumiDOVersion},
		"docker:version":                 auto.ConfigValue{Value: common.PulumiDockerVersion},
	}, nil
}

func (a *AwsExtensionProvider) Init(attributes map[string]interface{}) error {
	var err error

	region, ok := attributes["region"].(string)
	if !ok {
		return fmt.Errorf("missing region attribute")
	}

	a.Region = region

	a.config, err = ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	a.AwsConfig = &a.config.AwsConfig

	var isString bool

	iProject, hasProject := attributes["project"]
	a.ProjectName, isString = iProject.(string)
	if !hasProject || !isString || a.ProjectName == "" {
		// need a valid project name
		return fmt.Errorf("project is not set or invalid")
	}

	iStack, hasStack := attributes["stack"]
	a.StackName, isString = iStack.(string)
	if !hasStack || !isString || a.StackName == "" {
		// need a valid stack name
		return fmt.Errorf("stack is not set or invalid")
	}

	// Backwards compatible stack name
	// The existing providers in the CLI
	// Use the combined project and stack name
	a.FullStackName = fmt.Sprintf("%s-%s", a.ProjectName, a.StackName)

	return nil
}
