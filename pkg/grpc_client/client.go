package grpc_client

import (
	"fmt"

	"gitlab.udevs.io/delever/delever_user_service/config"
	"google.golang.org/grpc"

	pb "genproto/order_service"
)

type GrpcClientI interface {
	CustomerOrderService() pb.CustomerServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {
	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("order service dial host: %s port: %d",
			cfg.OrderServiceHost, cfg.OrderServicePort)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"order_service": pb.NewCustomerServiceClient(connOrder),
		},
	}, nil
}

func (g *GrpcClient) CustomerOrderService() pb.CustomerServiceClient {
	return g.connections["order_service"].(pb.CustomerServiceClient)
}
