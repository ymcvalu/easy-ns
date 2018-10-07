package main

import (
	"easy-ns/cgroups"
	"easy-ns/cgroups/subsystems"
	"easy-ns/container"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Run(tty bool, cmds []string, res *subsystems.ResourceConfig) {
	proc, pipe := container.NewProcess(tty)
	if proc == nil {
		log.Errorf("New proc error")
		return
	}

	if err := proc.Start(); err != nil {
		log.Errorf("when start proc error: %v", err)
		return
	}

	cgroupManager := cgroups.NewCgroupManager("lightDocker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(proc.Process.Pid)
	sendInitCommand(cmds, pipe)
	proc.Wait()
}

func sendInitCommand(cmds []string, pipe *os.File) {
	command := strings.Join(cmds, " ")
	log.Infof("command all is %s", command)
	pipe.WriteString(command)
	pipe.Close()
}
