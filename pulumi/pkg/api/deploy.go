package api

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Api struct {
	pulumi.ResourceState
	Name string
}

func NewApi(ctx pulumi.Context, name string, schedule deploypb.Resource_Api, opts ...pulumi.ResourceOption) (*Api, error) {
	res := &Api{Name: name}

	if err := ctx.RegisterComponentResource("custom:api:CustomApi", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Bucket

	return res, nil
}
