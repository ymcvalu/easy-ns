package main

import (
	"easy-ns/cgroups/subsystems"
	"easy-ns/container"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
		   light-docker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("miss container command")
		}
		var cmds = make([]string, 0, len(ctx.Args()))
		for _, arg := range ctx.Args() {
			cmds = append(cmds, arg)
		}
		tty := ctx.Bool("ti")
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: ctx.String("m"),
			CPUSet:      ctx.String("cpuset"),
			CPUShare:    ctx.String("cpushare"),
		}
		Run(tty, cmds, resConf)

		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(ctx *cli.Context) error {
		log.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}
