package controller

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (c *Controller) Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/", c.indexHandler)
	r.HandleFunc("/containers", c.scopeMiddleware(c.containersMiddleware(c.containersListHandler)))
	r.HandleFunc("/containers/{cID:.*}", c.scopeMiddleware(c.containerInspectMiddleware(c.containersInspectHandler)))
	r.HandleFunc("/services", c.scopeMiddleware(c.servicesListHandler))
	r.HandleFunc("/images", c.scopeMiddleware(c.imagesListHandler))
	r.HandleFunc("/networks", c.scopeMiddleware(c.networksListHandler))
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         c.config.ListenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.WithFields(logrus.Fields{
		"addr":  c.config.ListenAddr,
		"scope": c.config.Scope,
	}).Info("server started")
	return srv.ListenAndServe()
}

func (c *Controller) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("axiom metadata service\n"))
}
