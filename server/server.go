package server

import (
	"github.com/zipkero/grpc-server/config"
	"github.com/zipkero/grpc-server/grpc/client"
	"github.com/zipkero/grpc-server/grpc/server"
	"github.com/zipkero/grpc-server/router"
	"log"
	"net/http"
)

type Server struct {
	config    *config.Config
	router    *router.Router
	rpcClient *client.GRPCClient
	rpcServer *server.GRPCServer
}

func NewServer(configPath string) *Server {
	s := new(Server)
	s.config = config.NewConfig(configPath)
	s.router = router.NewRouter()
	s.rpcClient, _ = client.NewGRPCClient(s.config)
	s.rpcServer = server.NewGRPCServer(s.config)
	return s
}

func (s *Server) Start() error {
	s.router.AddRoute("GET", "/users", func(w http.ResponseWriter, req *http.Request) {
		if id, err := s.rpcClient.GetUser(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write([]byte(id))
		}
	})

	s.router.AddRoute("POST", "/users", func(w http.ResponseWriter, req *http.Request) {
		if id, err := s.rpcClient.GetUser(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write([]byte(id))
		}
	})

	go func() {
		if err := s.rpcServer.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServe(s.config.Server.Port, s.router); err != nil {
		return err
	}

	return nil
}
