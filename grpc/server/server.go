package server

import (
	"context"
	"github.com/zipkero/grpc-server/config"
	"github.com/zipkero/grpc-server/grpc/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	config *config.Config
	user.UserServiceServer
}

func NewGRPCServer(config *config.Config) *GRPCServer {
	return &GRPCServer{config: config}
}

func (s *GRPCServer) Start() error {
	if listener, err := net.Listen("tcp", s.config.Grpc.Port); err != nil {
		return err
	} else {
		rpcServer := grpc.NewServer([]grpc.ServerOption{}...)

		user.RegisterUserServiceServer(rpcServer, s)
		if err := rpcServer.Serve(listener); err != nil {
			return err
		}

		return nil
	}
}

func (s *GRPCServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Println("grpc call GetUser")
	return &user.GetUserResponse{}, nil
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log.Println("grpc call CreateUser")
	return &user.CreateUserResponse{}, nil
}
