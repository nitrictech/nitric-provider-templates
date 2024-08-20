package main

import (
	_ "embed"

	"github.com/nitrictech/nitric-provider-template/custom-provider/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
)

// Embed the runtime provider binary
//
//go:embed runtime-bin
var runtimeBin []byte

var runtimeProvider = func() []byte {
	return runtimeBin
}

// Start the deployment server
func main() {
	stack := deploy.NewNitricCustomPulumiProvider()

	providerServer := provider.NewPulumiProviderServer(stack, runtimeProvider)

	providerServer.Start()
}
