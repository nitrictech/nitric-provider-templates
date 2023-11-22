package websocket

import (
	deploypb "github.com/nitrictech/nitric/core/pkg/proto/deploy/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Websocket struct {
	pulumi.ResourceState
	Name string
}

func NewWebsocket(ctx pulumi.Context, name string, schedule deploypb.Resource_Websocket, opts ...pulumi.ResourceOption) (*Websocket, error) {
	res := &Websocket{Name: name}

	if err := ctx.RegisterComponentResource("custom:websocket:CustomWebsocket", name, res, opts...); err != nil {
		return nil, err
	}

	// TODO: Implement Websockets

	return res, nil
}
