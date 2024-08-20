package websocket

import (
	"context"

	websocketpb "github.com/nitrictech/nitric/core/pkg/proto/websockets/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WebsocketServer struct{}

var _ websocketpb.WebsocketServer = &WebsocketServer{}

// Get the details of the websocket
func (*WebsocketServer) SocketDetails(context.Context, *websocketpb.WebsocketDetailsRequest) (*websocketpb.WebsocketDetailsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Send a messages to a websocket
func (*WebsocketServer) SendMessage(context.Context, *websocketpb.WebsocketSendRequest) (*websocketpb.WebsocketSendResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

// Close a websocket connection
// This can be used to force a client to disconnect
func (*WebsocketServer) CloseConnection(context.Context, *websocketpb.WebsocketCloseConnectionRequest) (*websocketpb.WebsocketCloseConnectionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func New() (*WebsocketServer, error) {
	return &WebsocketServer{}, nil
}
