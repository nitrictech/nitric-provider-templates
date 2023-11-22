package topic

import (
	"context"

	topicpb "github.com/nitrictech/nitric/core/pkg/proto/topic/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TopicServer struct{}

var _ topicpb.TopicServiceServer = &TopicServer{}

// Updates a secret, creating a new one if it doesn't already exist
func (srv *TopicServer) Publish(ctx context.Context, req *topicpb.TopicPublishRequest) (*topicpb.TopicPublishResponse, error) {
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func NewServer() *TopicServer {
	return &TopicServer{}
}
