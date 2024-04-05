package topic

import (
	"context"

	"github.com/nitrictech/nitric-provider-template/custom-provider/pkg/runtime/resource"
	topicpb "github.com/nitrictech/nitric/core/pkg/proto/topics/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TopicsServer struct{}

var _ topicpb.TopicsServer = &TopicsServer{}

// Updates a secret, creating a new one if it doesn't already exist
func (srv *TopicsServer) Publish(ctx context.Context, req *topicpb.TopicPublishRequest) (*topicpb.TopicPublishResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func New(provider *resource.ResourceServer) (*TopicsServer, error) {
	return &TopicsServer{}, nil
}
