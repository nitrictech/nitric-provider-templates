package main

import (
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime"
	"github.com/nitrictech/nitric/core/pkg/logger"
	"github.com/nitrictech/nitric/core/pkg/server"
)

func main() {
	// Perform any custom initialization here, such as setting up a custom resolver

	// Create the runtime server and run it
	m, err := runtime.NewRuntimeServer()
	if err != nil {
		logger.Fatalf("there was an error initializing the runtime server: %v", err)
	}

	server.Run(m)
}
