package config

type Config struct {
	Server struct {
		Host string
	}
	GRPC struct {
		Port int
	}
}
