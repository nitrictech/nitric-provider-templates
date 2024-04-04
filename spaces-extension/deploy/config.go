package deploy

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/nitrictech/nitric/cloud/aws/deploy"
)

type ExtensionConfig struct {
	deploy.AwsConfig
	Token  string        `mapstructure:"token"`
	Spaces *SpacesConfig `mapstructure:"spaces"`
}

type SpacesConfig struct {
	Key    string `mapstructure:"key"`
	Secret string `mapstructure:"secret"`
	Region string `mapstructure:"region"`
}

func ConfigFromAttributes(attributes map[string]interface{}) (*ExtensionConfig, error) {
	extensionConfig := &ExtensionConfig{}
	err := mapstructure.Decode(attributes, extensionConfig)
	if err != nil {
		return nil, err
	}

	if extensionConfig.Token == "" || extensionConfig.Spaces == nil || extensionConfig.Spaces.Key == "" || extensionConfig.Spaces.Secret == "" {
		return nil, fmt.Errorf("invalid config: require a digital ocean token, spaces access key and access secret")
	}

	awsConfig, err := deploy.ConfigFromAttributes(attributes)
	if err != nil {
		return nil, err
	}

	extensionConfig.AwsConfig = *awsConfig

	return extensionConfig, nil
}
