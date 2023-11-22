package exec

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This represents a unique unit of execution at the moment this is a container but could also be many things e.g. WASM, Binary, source zip etc.
type Exec struct {
	pulumi.ResourceState
	Name string
}

func NewHttp(ctx pulumi.Context, name string, schedule deploypb.Resource_Api, opts ...pulumi.ResourceOption) (*Exec, error) {
	res := &Exec{Name: name}

	if err := ctx.RegisterComponentResource("custom:exec:CustomExec", name, res, opts...); err != nil {
		return nil, err
	}

	return res, nil
}
