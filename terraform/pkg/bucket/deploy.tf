package bucket

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Bucket struct {
	pulumi.ResourceState
	Name string
}

func NewCollection(ctx pulumi.Context, name string, schedule deploypb.Resource_Bucket, opts ...pulumi.ResourceOption) (*Bucket, error) {
	res := &Bucket{Name: name}

	if err := ctx.RegisterComponentResource("custom:storage:CustomBucket", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Bucket

	return res, nil
}
