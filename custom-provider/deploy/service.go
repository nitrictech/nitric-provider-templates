package deploy

import (
	"fmt"

	"github.com/nitrictech/nitric/cloud/common/deploy/image"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/common/deploy/pulumix"
	"github.com/nitrictech/nitric/cloud/common/deploy/resources"
	"github.com/nitrictech/nitric/cloud/common/deploy/tags"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ecr"
	awslambda "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createEcrRepository(ctx *pulumi.Context, parent pulumi.Resource, stackId string, name string) (*ecr.Repository, error) {
	return ecr.NewRepository(ctx, name, &ecr.RepositoryArgs{
		ForceDelete: pulumi.BoolPtr(true),
		Tags:        pulumi.ToStringMap(tags.Tags(stackId, name, resources.Service)),
	}, pulumi.Parent(parent))
}

func createImage(ctx *pulumi.Context, parent pulumi.Resource, name string, authToken *ecr.GetAuthorizationTokenResult, repo *ecr.Repository, config *pulumix.NitricPulumiServiceConfig, runtime provider.RuntimeProvider) (*image.Image, error) {
	// ensure valid image configuriation
	if config.GetImage() == nil {
		return nil, fmt.Errorf("aws provider can only deploy service with an image source")
	}

	if config.GetImage().GetUri() == "" {
		return nil, fmt.Errorf("aws provider can only deploy service with an image source")
	}

	if config.Type == "" {
		config.Type = "default"
	}

	// create the image
	return image.NewImage(ctx, name, &image.ImageArgs{
		SourceImage:   config.GetImage().GetUri(),
		RepositoryUrl: repo.RepositoryUrl,
		Server:        pulumi.String(authToken.ProxyEndpoint),
		Username:      pulumi.String(authToken.UserName),
		Password:      pulumi.String(authToken.Password),
		Runtime:       runtime(),
	}, pulumi.Parent(parent), pulumi.DependsOn([]pulumi.Resource{repo}))
}

func (a *NitricAwsPulumiProvider) Service(ctx *pulumi.Context, parent pulumi.Resource, name string, config *pulumix.NitricPulumiServiceConfig, runtime provider.RuntimeProvider) error {
	opts := []pulumi.ResourceOption{pulumi.Parent(parent)}

	// Create the ECR repository to push the image to
	repo, err := createEcrRepository(ctx, parent, a.StackId, name)
	if err != nil {
		return err
	}

	// Create the image
	image, err := createImage(ctx, parent, name, a.EcrAuthToken, repo, config, runtime)
	if err != nil {
		return err
	}

	// Create the Lambda Function
	a.Lambdas[name], err = awslambda.NewFunction(ctx, name, &awslambda.FunctionArgs{
		ImageUri:    image.URI(),
		PackageType: pulumi.String("Image"),
		Environment: awslambda.FunctionEnvironmentArgs{Variables: envVars},
	}, append([]pulumi.ResourceOption{pulumi.DependsOn([]pulumi.Resource{image})}, opts...)...)
	if err != nil {
		return err
	}

	return nil
}
