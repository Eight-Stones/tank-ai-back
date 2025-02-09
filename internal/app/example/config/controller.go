package config

import (
	"go-micro.dev/v4/config"
)

const (
	defaultHost = "localhost"
	defaultPort = 8080
)

type Controller struct {
	RestBook RestBook
}

type RestBook struct {
	Host string
	Port int
}

func newController(cfg config.Config) Controller {
	return Controller{
		RestBook: newRestBook(cfg),
	}
}

func newRestBook(cfg config.Config) RestBook {
	return RestBook{
		Host: cfg.Get("controller", "rest", "host").String(defaultHost),
		Port: cfg.Get("controller", "rest", "port").Int(defaultPort),
	}
}
