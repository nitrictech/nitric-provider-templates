package websocket

import (
	"context"

	websocketpb "github.com/nitrictech/nitric/core/pkg/proto/websocket/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WebsocketServer struct{}

var _ websocketpb.WebsocketServiceServer = &WebsocketServer{}

// Send a messages to a websocket
func (*WebsocketServer) Send(context.Context, *websocketpb.WebsocketSendRequest) (*websocketpb.WebsocketSendResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Close a websocket connection
// This can be used to force a client to disconnect
func (*WebsocketServer) Close(context.Context, *websocketpb.WebsocketCloseRequest) (*websocketpb.WebsocketCloseResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func NewServer() *WebsocketServer {
	return &WebsocketServer{}
}
