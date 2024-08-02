package client

import (
	"github.com/zipkero/grpc-server/config"
	"github.com/zipkero/grpc-server/grpc/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	config     *config.Config
	userClient user.UserServiceClient
}

func NewGRPCClient(config *config.Config) *GRPCClient {
	c := &GRPCClient{
		config: config,
	}
	return c
}

func (c *GRPCClient) Start() error {
	conn, err := grpc.NewClient(c.config.Grpc.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	c.userClient = user.NewUserServiceClient(conn)

	return nil
}
