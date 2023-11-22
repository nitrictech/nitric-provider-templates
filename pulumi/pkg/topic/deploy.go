package topic

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Topic struct {
	pulumi.ResourceState
	Name string
}

func NewTopic(ctx pulumi.Context, name string, topic deploypb.Resource_Topic, opts ...pulumi.ResourceOption) (*Topic, error) {
	res := &Topic{Name: name}

	if err := ctx.RegisterComponentResource("custom:topic:CustomTopic", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Topic

	return res, nil
}
