package collection

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Collection struct {
	pulumi.ResourceState
	Name string
}

func NewCollection(ctx pulumi.Context, name string, schedule deploypb.Resource_Collection, opts ...pulumi.ResourceOption) (*Collection, error) {
	res := &Collection{Name: name}

	if err := ctx.RegisterComponentResource("custom:collection:CustomCollection", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Collection

	return res, nil
}
