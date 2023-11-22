package secret

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Secret struct {
	pulumi.ResourceState
	Name string
}

func NewSecret(ctx pulumi.Context, name string, secret deploypb.Resource_Secret, opts ...pulumi.ResourceOption) (*Secret, error) {
	res := &Secret{Name: name}

	if err := ctx.RegisterComponentResource("custom:secret:CustomSecret", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Secret

	return res, nil
}
