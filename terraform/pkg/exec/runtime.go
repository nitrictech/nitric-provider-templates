package exec

import "github.com/nitrictech/nitric/core/pkg/worker/pool"

type GatewayService struct {
}

// Start the Gateway
func (*GatewayService) Start(pool pool.WorkerPool) error {
	// Start up your functions ingress here
	// This could be anything from a simple HTTP server to an AWS Lambda queue processor or a function runtime host (e.g. OpenFaaS, Knative Serving, OpenWhisk etc.)
}

// Stop the Gateway
func (*GatewayService) Stop() error {

}
