package main

import (
	"github.com/zipkero/grpc-server/server"
	"path/filepath"
)

func main() {
	s := server.NewServer(filepath.Join("..", "config.toml"))
	if err := s.Start(); err != nil {
		panic(err)
	}
}
