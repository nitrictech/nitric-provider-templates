package keyvalue

import (
	"context"

	"github.com/nitrictech/nitric-provider-template/custom-provider/runtime/resource"
	queuespb "github.com/nitrictech/nitric/core/pkg/proto/queues/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueuesServer struct{}

var _ queuespb.QueuesServer = &QueuesServer{}

// Complete implements queuespb.QueuesServer.
func (*QueuesServer) Complete(context.Context, *queuespb.QueueCompleteRequest) (*queuespb.QueueCompleteResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Dequeue implements queuespb.QueuesServer.
func (*QueuesServer) Dequeue(context.Context, *queuespb.QueueDequeueRequest) (*queuespb.QueueDequeueResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

// Enqueue implements queuespb.QueuesServer.
func (*QueuesServer) Enqueue(context.Context, *queuespb.QueueEnqueueRequest) (*queuespb.QueueEnqueueResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func New(provider *resource.ResourceServer) (*QueuesServer, error) {
	return &QueuesServer{}, nil
}
