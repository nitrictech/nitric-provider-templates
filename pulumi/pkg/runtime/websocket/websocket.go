package websocket

import (
	"context"

	"github.com/nitrictech/nitric-provider-template/pulumi/pkg/runtime/resource"
	websocketpb "github.com/nitrictech/nitric/core/pkg/proto/websockets/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WebsocketServer struct{}

var _ websocketpb.WebsocketServer = &WebsocketServer{}

// Get the details of the websocket
func (*WebsocketServer) SocketDetails(context.Context, *websocketpb.WebsocketDetailsRequest) (*websocketpb.WebsocketDetailsResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Send a messages to a websocket
func (*WebsocketServer) SendMessage(context.Context, *websocketpb.WebsocketSendRequest) (*websocketpb.WebsocketSendResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Close a websocket connection
// This can be used to force a client to disconnect
func (*WebsocketServer) CloseConnection(context.Context, *websocketpb.WebsocketCloseConnectionRequest) (*websocketpb.WebsocketCloseConnectionResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func New(provider *resource.ResourceServer) (*WebsocketServer, error) {
	return &WebsocketServer{}, nil
}
