package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	ctx "github.com/gorilla/context"
	"github.com/sirupsen/logrus"
)

func (c *Controller) containersListHandler(w http.ResponseWriter, r *http.Request) {
	cnts, ok := ctx.GetOk(r, ContainersKey)
	if !ok {
		msg := "unable to get containers"
		logrus.Error(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	containers := cnts.([]types.Container)
	if err := json.NewEncoder(w).Encode(containers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) containersInspectHandler(w http.ResponseWriter, r *http.Request) {
	cnt, ok := ctx.GetOk(r, ContainerInspectKey)
	if !ok {
		msg := "unable to inspect container"
		logrus.Error(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	container := cnt.(types.ContainerJSON)
	if err := json.NewEncoder(w).Encode(container); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getContainerFromIP(ip string) (*types.Container, error) {
	// get controller network
	network, err := c.getControllerNetwork()
	if err != nil {
		return nil, err
	}

	// get all containers and filter on network
	args := filters.NewArgs()
	args.Add("network", network)

	containers, err := c.client.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: args,
	})
	if err != nil {
		return nil, err
	}

	for _, cnt := range containers {
		for _, n := range cnt.NetworkSettings.Networks {
			if n.IPAddress == ip {
				return &cnt, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to find container")
}
