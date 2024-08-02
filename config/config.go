package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	Grpc struct {
		Port string
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := toml.NewDecoder(file).Decode(c); err != nil {
		log.Fatal(err)
	}

	return c
}
