package schedule

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schedule struct {
	pulumi.ResourceState
	Name string
}

func NewSchedule(ctx pulumi.Context, name string, schedule deploypb.Resource_Schedule, opts ...pulumi.ResourceOption) (*Schedule, error) {
	res := &Schedule{Name: name}

	if err := ctx.RegisterComponentResource("custom:secret:CustomSchedule", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Schedule

	return res, nil
}
