package policy

import (
	resourcepb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Policy struct {
	pulumi.ResourceState
	Name string
}

func NewPolicy(ctx pulumi.Context, name string, schedule resourcepb.Resource_Policy, opts ...pulumi.ResourceOption) (*Policy, error) {
	res := &Policy{Name: name}

	if err := ctx.RegisterComponentResource("custom:policy:CustomPolicy", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Policy

	return res, nil
}
