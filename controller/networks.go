package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
)

func (c *Controller) networksListHandler(w http.ResponseWriter, r *http.Request) {
	networks, err := c.client.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(networks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
