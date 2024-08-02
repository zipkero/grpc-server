package client

import (
	"context"
	"github.com/zipkero/grpc-server/config"
	"github.com/zipkero/grpc-server/grpc/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	config     *config.Config
	userClient user.UserServiceClient
}

func NewGRPCClient(config *config.Config) (*GRPCClient, error) {
	c := &GRPCClient{
		config: config,
	}

	if conn, err := grpc.NewClient(c.config.Grpc.Port, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.userClient = user.NewUserServiceClient(conn)
	}

	return c, nil
}

func (c *GRPCClient) GetUser() (string, error) {
	if res, err := c.userClient.GetUser(context.Background(), &user.GetUserRequest{}); err != nil {
		return "", err
	} else {
		return res.Id, nil
	}
}

func (c *GRPCClient) CreateUser() (string, error) {
	if res, err := c.userClient.CreateUser(context.Background(), &user.CreateUserRequest{}); err != nil {
		return "", err
	} else {
		return res.Id, nil
	}

}
