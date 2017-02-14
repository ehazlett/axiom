package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
)

func (c *Controller) servicesListHandler(w http.ResponseWriter, r *http.Request) {
	services, err := c.client.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(services); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
