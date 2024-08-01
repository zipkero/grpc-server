package server

import "grpc-server/router"

type Server struct {
	router *router.Router
}

func NewServer() *Server {
	s := new(Server)
	s.router = router.NewRouter()
	return s
}

func (s *Server) Start() {

}
