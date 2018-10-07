package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = `light-docker is a simple container runtime implementation.
			   The purpose of this projects is to learn how docker works and how to write a docker by ourselves.
			   Enjoy it, just for fun.`

func main() {
	app := cli.NewApp()
	app.Name = "light-docker"
	app.Usage = usage

	app.Commands = []cli.Command{
		runCommand,
		initCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
