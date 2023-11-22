package http

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Http struct {
	pulumi.ResourceState
	Name string
}

func NewHttp(ctx pulumi.Context, name string, schedule deploypb.Resource_Api, opts ...pulumi.ResourceOption) (*Http, error) {
	res := &Http{Name: name}

	if err := ctx.RegisterComponentResource("custom:http:CustomHttp", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Http

	return res, nil
}
