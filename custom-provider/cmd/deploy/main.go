package main

import (
	"github.com/nitrictech/nitric-provider-template/custom-provider/deploy"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
)

// Embed the runtime provider
//
//go:embed runtime-extension-aws
var runtimeBin []byte

var runtimeProvider = func() []byte {
	return runtimeBin
}

// Start the deployment server
func main() {
	stack := deploy.NewNitricCustomPulumiProvider()

	providerServer := provider.NewPulumiProviderServer(stack, runtime.NitricCustomRuntime)

	providerServer.Start()
}
