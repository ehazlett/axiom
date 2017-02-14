package controller

import (
	"context"
	"net"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	ctx "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (c *Controller) scopeMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"middleware": "scope",
			}).Errorf("unable to get host and port from remote address: %s", err)
			fn(w, r)
			return
		}

		cnt, err := c.getContainerFromIP(host)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"middleware": "scope",
			}).Errorf("unable to get container from request")
			fn(w, r)
			return
		}

		logrus.WithFields(logrus.Fields{
			"container": cnt.ID,
			"remote":    r.RemoteAddr,
			"path":      r.URL,
		}).Debug("request")

		ctx.Set(r, CurrentContainerKey, cnt)

		fn(w, r)
	}
}

func (c *Controller) containersMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cnt, ok := ctx.GetOk(r, CurrentContainerKey)
		if !ok {
			logrus.WithFields(logrus.Fields{
				"middleware": "containers",
			}).Warnf("unable to get container id from context")
		}
		args := filters.NewArgs()
		if c.config.Scope != GlobalScope {
			container := cnt.(*types.Container)
			args.Add("id", container.ID)
		}

		containers, err := c.client.ContainerList(context.Background(), types.ContainerListOptions{
			Filters: args,
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"middleware": "containers",
			}).Errorf("unable to get containers: %s", err)
			fn(w, r)
			return
		}

		ctx.Set(r, ContainersKey, containers)

		fn(w, r)
	}
}

func (c *Controller) containerInspectMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["cID"]
		if id == "" {
			http.Error(w, "you must specify an id", http.StatusBadRequest)
			return
		}

		if c.config.Scope != GlobalScope {
			cnt, ok := ctx.GetOk(r, CurrentContainerKey)
			if !ok {
				logrus.WithFields(logrus.Fields{
					"middleware": "containerInspect",
				}).Warnf("unable to get container id from context")
			}

			container := cnt.(*types.Container)
			// TODO: user filters to get partial ID
			if container.ID != id {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
		}

		container, err := c.client.ContainerInspect(context.Background(), id)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"middleware": "containerInspect",
			}).Errorf("unable to inspect container: %s", err)
			fn(w, r)
			return
		}

		ctx.Set(r, ContainerInspectKey, container)

		fn(w, r)
	}
}
