package deploy

import (
	"github.com/nitrictech/nitric/cloud/aws/deploy"
	common "github.com/nitrictech/nitric/cloud/common/deploy"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AwsExtensionProvider struct {
	deploy.NitricAwsPulumiProvider

	*common.CommonStackDetails

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
	config, err := a.NitricAwsPulumiProvider.Config()
	if err != nil {
		return nil, err
	}

	config["digitalocean:token"] = auto.ConfigValue{Value: a.config.Token, Secret: true}
	config["digitalocean:spaces_access_id"] = auto.ConfigValue{Value: a.config.Spaces.Key, Secret: true}
	config["digitalocean:spaces_secret_key"] = auto.ConfigValue{Value: a.config.Spaces.Secret, Secret: true}
	config["digitalocean:version"] = auto.ConfigValue{Value: "4.27.0"}

	return config, nil
}

func (a *AwsExtensionProvider) Init(attributes map[string]interface{}) error {
	var err error

	a.CommonStackDetails, err = common.CommonStackDetailsFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	a.config, err = ConfigFromAttributes(attributes)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Bad stack configuration: %s", err)
	}

	a.AwsConfig = &a.config.AwsConfig

	return nil
}
