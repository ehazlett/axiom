package controller

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
)

func (c *Controller) getControllerContainer() (*types.ContainerJSON, error) {
	id, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	cnt, err := c.client.ContainerInspect(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &cnt, nil
}

func (c *Controller) getControllerNetwork() (string, error) {
	cnt, err := c.getControllerContainer()
	if err != nil {
		return "", err
	}

	// TODO: support multiple networks
	for k, _ := range cnt.NetworkSettings.Networks {
		return k, nil
	}

	return "", fmt.Errorf("controller network not found")
}
