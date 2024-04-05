package http

import (
	"github.com/nitrictech/nitric-provider-template/custom-provider/pkg/runtime/resource"
	"github.com/nitrictech/nitric/core/pkg/gateway"
)

func NewHttpGateway(provider *resource.ResourceServer) (gateway.GatewayService, error) {
	return &gateway.UnimplementedGatewayPlugin{}, nil
}
