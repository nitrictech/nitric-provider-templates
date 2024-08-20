package runtime

import (
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/api"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/gateway"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/keyvalue"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/queue"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/secret"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/sql"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/storage"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/topic"
	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/websocket"
	"github.com/nitrictech/nitric/core/pkg/server"
)

func NewRuntimeServer(opts ...server.ServerOption) (*server.NitricServer, error) {

	// Create the default plugins
	// this is a good place to pass in other useful values, such as a runtime cloud resource resolver
	apiPlugin, _ := api.New()
	gatewayPlugin, _ := gateway.New()
	keyValuePlugin, _ := keyvalue.New()
	queuesPlugin, _ := queue.New()
	secretPlugin, _ := secret.New()
	storagePlugin, _ := storage.New()
	sqlPlugin, _ := sql.New()
	topicsPlugin, _ := topic.New()
	websocketPlugin, _ := websocket.New()

	// Set the default options for your runtime server, extensions may override these.
	defaultOptions := []server.ServerOption{
		server.WithApiPlugin(apiPlugin),
		server.WithGatewayPlugin(gatewayPlugin),
		server.WithKeyValuePlugin(keyValuePlugin),
		server.WithSecretManagerPlugin(secretPlugin),
		server.WithStoragePlugin(storagePlugin),
		server.WithWebsocketPlugin(websocketPlugin),
		server.WithTopicsPlugin(topicsPlugin),
		server.WithQueuesPlugin(queuesPlugin),
		server.WithSqlPlugin(sqlPlugin),
	}

	// Merge the default options with the provided options, the provided options will override the defaults.
	options := append(defaultOptions, opts...)

	return server.New(options...)
}
