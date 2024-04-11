package deploy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
)

type DeploymentServer struct {
	deploymentspb.UnimplementedDeploymentServer
}

func (s *DeploymentServer) Up(req *deploymentspb.DeploymentUpRequest, srv deploymentspb.Deployment_UpServer) error {
	attributes := req.Attributes.AsMap()

	output, ok := attributes["output"].(string)
	if !ok {
		output = "./output"
	}

	outputHcl, ok := attributes["hcl"].(bool)
	if !ok {
		outputHcl = false
	}

	app := cdktf.NewApp(&cdktf.AppConfig{
		HclOutput: &outputHcl,
		Outdir:    &output,
	})

	NewMyStack(app, "cdktf", req)

	app.Synth()

	return nil
}

func NewDeploymentServer() *DeploymentServer {
	return &DeploymentServer{}
}
