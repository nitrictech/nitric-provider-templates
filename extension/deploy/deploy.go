package deploy

import (
	"github.com/nitrictech/nitric/cloud/aws/deploy"
)

type CustomAwsExtensionProvider struct {
	*deploy.NitricAwsPulumiProvider
}

func NewCustomAwsExtensionProvider() *CustomAwsExtensionProvider {
	return &CustomAwsExtensionProvider{
		NitricAwsPulumiProvider: deploy.NewNitricAwsProvider(),
	}
}
