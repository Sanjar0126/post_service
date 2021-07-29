package grpc_client

import (
	"fmt"

	"github.com/Sanjar0126/post_service/config"
	"google.golang.org/grpc"

	pb "genproto/post_service"
)

type GrpcClientI interface {
	CustomerOrderService() pb.PostServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d",
			cfg.PostServiceHost, cfg.PostServicePort)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"post_service": pb.NewPostServiceClient(connOrder),
		},
	}, nil
}

func (g *GrpcClient) CustomerOrderService() pb.PostServiceClient {
	return g.connections["post_service"].(pb.PostServiceClient)
}
