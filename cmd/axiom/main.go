package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/ehazlett/axiom/version"
	"github.com/sirupsen/logrus"
)

func main() {
	app := cli.NewApp()
	app.Name = version.Name()
	app.Usage = version.Description()
	app.Version = version.Version()
	app.Author = "@ehazlett"
	app.Email = ""
	app.Before = func(c *cli.Context) error {
		// enable debug
		if c.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.Debug("debug enabled")
		}

		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
		cli.StringFlag{
			Name:  "listen, l",
			Usage: "listen address",
			Value: ":80",
		},
		cli.StringFlag{
			Name:  "scope, s",
			Usage: "metadata access scope (global, limited)",
			Value: "global",
		},
	}
	app.Action = func(c *cli.Context) error {
		logrus.Info(version.FullVersion())

		ctrl, err := getController(c)
		if err != nil {
			return err
		}

		if err := ctrl.Run(); err != nil {
			return err
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

}
