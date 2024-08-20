package gateway

import (
	"fmt"

	"github.com/nitrictech/nitric/core/pkg/gateway"
)

type GatewayServer struct{}

var _ gateway.GatewayService = &GatewayServer{}

func (g *GatewayServer) Start(opts *gateway.GatewayStartOpts) error {
	return fmt.Errorf("Gateway Start is Unimplemented")
}

func (g *GatewayServer) Stop() error {
	return fmt.Errorf("Gateway Stop is Unimplemented")
}

func New() (*GatewayServer, error) {
	return &GatewayServer{}, nil
}
