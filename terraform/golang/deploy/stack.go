package deploy

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	terraformApi "github.com/nitrictech/nitric-provider-templates/terraform/golang/generated/api"
	terraformService "github.com/nitrictech/nitric-provider-templates/terraform/golang/generated/cloudrun"
	terraformPolicies "github.com/nitrictech/nitric-provider-templates/terraform/golang/generated/policies"
	terraformStorage "github.com/nitrictech/nitric-provider-templates/terraform/golang/generated/storage"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

func NewMyStack(scope constructs.Construct, id string, req *deploymentspb.DeploymentUpRequest) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Deploy common policies
	policies := terraformPolicies.NewPolicies(stack, jsii.String("policies"), &terraformPolicies.PoliciesConfig{
		ProjectId: jsii.String("TODO"),
	})

	// Deploy the services
	for _, resource := range req.Spec.Resources {
		switch resource.Config.(type) {
		case *deploymentspb.Resource_Service:
			terraformService.NewCloudrun(stack, jsii.String(resource.Id.Name), &terraformService.CloudrunConfig{
				ServiceName: jsii.String(resource.Id.Name),
			})
		}
	}

	// Deploy the Buckets
	for _, resource := range req.Spec.Resources {
		switch resource.Config.(type) {
		case *deploymentspb.Resource_Bucket:
			terraformStorage.NewStorage(stack, jsii.String(resource.Id.Name), &terraformStorage.StorageConfig{
				BucketName: jsii.String(resource.Id.Name),
			})
		}
	}

	// Deploy the APIs
	for _, resource := range req.Spec.Resources {
		switch resource.Config.(type) {
		case *deploymentspb.Resource_Api:
			// Transform the open api spec into a GCP ApiGateway compatible spec
			terraformApi.NewApi(stack, jsii.String(resource.Id.Name), &terraformApi.ApiConfig{
				ApiName: jsii.String(resource.Id.Name),
			})
		}
	}

	// Deploy the Policies
	for _, resource := range req.Spec.Resources {
		switch resource.Config.(type) {
		case *deploymentspb.Resource_Policy:

		}
	}

	terraformApi.NewApi(stack, jsii.String("api"), &terraformApi.ApiConfig{})

	// The code that defines your stack goes here

	return stack
}
