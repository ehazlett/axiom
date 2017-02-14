package controller

import (
	"github.com/docker/docker/client"
)

type Controller struct {
	config *Config
	client *client.Client
}

func NewController(cfg *Config) (*Controller, error) {
	c, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Controller{
		config: cfg,
		client: c,
	}, nil
}
