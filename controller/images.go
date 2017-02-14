package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
)

func (c *Controller) imagesListHandler(w http.ResponseWriter, r *http.Request) {
	images, err := c.client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(images); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
