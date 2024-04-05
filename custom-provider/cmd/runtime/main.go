package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/http"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/keyvalue"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/resource"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/secret"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/storage"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/topic"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/websocket"
	"github.com/nitrictech/nitric/core/pkg/logger"
	"github.com/nitrictech/nitric/core/pkg/membrane"
)

func main() {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt, syscall.SIGINT)

	membraneOpts := membrane.DefaultMembraneOptions()

	provider, err := resource.New()
	if err != nil {
		logger.Fatalf("could not create aws provider: %v", err)
		return
	}

	// Load the appropriate gateway based on the environment.
	membraneOpts.GatewayPlugin, _ = http.NewHttpGateway(nil)

	membraneOpts.SecretManagerPlugin, _ = secret.New(provider)
	membraneOpts.KeyValuePlugin, _ = keyvalue.New(provider)
	membraneOpts.TopicsPlugin, _ = topic.New(provider)
	membraneOpts.StoragePlugin, _ = storage.New(provider)
	membraneOpts.ResourcesPlugin = provider
	membraneOpts.WebsocketPlugin, _ = websocket.New(provider)

	m, err := membrane.New(membraneOpts)
	if err != nil {
		logger.Fatalf("There was an error initializing the membrane server: %v", err)
	}

	errChan := make(chan error)
	// Start the Membrane server
	go func(chan error) {
		errChan <- m.Start()
	}(errChan)

	select {
	case membraneError := <-errChan:
		logger.Errorf("Membrane Error: %v, exiting\n", membraneError)
	case sigTerm := <-term:
		logger.Debugf("Received %v, exiting\n", sigTerm)
	}

	m.Stop()
}
