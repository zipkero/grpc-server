package client

import "github.com/zipkero/grpc-server/config"

type GRPCClient struct {
}

func NewGRPCClient(config config.Config) *GRPCClient {
	c := &GRPCClient{}
	return c
}
