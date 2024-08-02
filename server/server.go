package server

import (
	"grpc-server/config"
	"grpc-server/router"
	"net/http"
)

type Server struct {
	config *config.Config
	router *router.Router
}

func NewServer(configPath string) *Server {
	s := new(Server)
	s.config = config.NewConfig(configPath)
	s.router = router.NewRouter()
	return s
}

func (s *Server) Start() error {
	s.router.AddRoute("GET", "/", func(w http.ResponseWriter, req *http.Request) {
		if _, err := w.Write([]byte("grpc server")); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	return http.ListenAndServe(s.config.Server.Port, s.router)
}
