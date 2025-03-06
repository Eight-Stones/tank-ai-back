package config

import (
	"go-micro.dev/v4/config"
)

const (
	defaultHost     = "localhost"
	defaultRestPort = 8080
	defaultGrpcPort = 50051
)

type Controller struct {
	RestTank Rest
	GrpcTank Grpc
}

type Rest struct {
	Host string
	Port int
}

type Grpc struct {
	Host string
	Port int
}

func newController(cfg config.Config) Controller {
	return Controller{
		RestTank: newTankBook(cfg),
		GrpcTank: newGrpcBook(cfg),
	}
}

func newTankBook(cfg config.Config) Rest {
	return Rest{
		Host: cfg.Get("controller", "rest", "host").String(defaultHost),
		Port: cfg.Get("controller", "rest", "port").Int(defaultRestPort),
	}
}

func newGrpcBook(cfg config.Config) Grpc {
	return Grpc{
		Host: cfg.Get("controller", "grpc", "host").String(defaultHost),
		Port: cfg.Get("controller", "grpc", "port").Int(defaultGrpcPort),
	}
}
