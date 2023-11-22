package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eapache/queue"
	"github.com/nitrictech/nitric-provider-template/pulumi/pkg/topic"
	"github.com/nitrictech/nitric/cloud/aws/runtime/core"
	"github.com/nitrictech/nitric/cloud/aws/runtime/websocket"
	"github.com/nitrictech/nitric/core/pkg/membrane"
	"github.com/nitrictech/nitric/core/pkg/utils"
	"github.com/viant/toolbox/secret"
	"google.golang.org/api/storage/v1"
)

func main() {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	signal.Notify(term, os.Interrupt, syscall.SIGINT)

	gatewayEnv := utils.GetEnv("GATEWAY_ENVIRONMENT", "lambda")

	membraneOpts := membrane.DefaultMembraneOptions()

	provider, err := core.New()
	if err != nil {
		log.Fatalf("could not create aws provider: %v", err)
		return
	}

	// Load the appropriate gateway based on the environment.
	switch gatewayEnv {
	case "lambda":
		membraneOpts.GatewayPlugin, _ = lambda_service.New(provider)
	default:
		membraneOpts.GatewayPlugin, _ = base_http.New(nil)
	}

	membraneOpts.SecretPlugin, _ = secret.NewServer()
	membraneOpts.DocumentPlugin, _ = document.NewServer()
	membraneOpts.EventsPlugin, _ = topic.NewServer()
	membraneOpts.QueuePlugin, _ = queue.NewServer()
	membraneOpts.StoragePlugin = storage.NewServer()
	// membraneOpts.ResourcesPlugin = provider
	membraneOpts.CreateTracerProvider = newTracerProvider
	membraneOpts.WebsocketPlugin = websocket.NewServer()

	m, err := membrane.New(membraneOpts)
	if err != nil {
		log.Default().Fatalf("There was an error initialising the membrane server: %v", err)
	}

	errChan := make(chan error)
	// Start the Membrane server
	go func(chan error) {
		errChan <- m.Start()
	}(errChan)

	select {
	case membraneError := <-errChan:
		log.Default().Printf("Membrane Error: %v, exiting\n", membraneError)
	case sigTerm := <-term:
		log.Default().Printf("Received %v, exiting\n", sigTerm)
	}

	m.Stop()
}
