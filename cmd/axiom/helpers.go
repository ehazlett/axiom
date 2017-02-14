package main

import (
	"github.com/codegangsta/cli"
	"github.com/ehazlett/axiom/controller"
)

func getController(c *cli.Context) (*controller.Controller, error) {
	cfg := &controller.Config{
		ListenAddr: c.GlobalString("listen"),
	}

	switch c.GlobalString("scope") {
	case "global":
		cfg.Scope = controller.GlobalScope
	case "limited":
		cfg.Scope = controller.LimitedScope
	}

	return controller.NewController(cfg)
}
