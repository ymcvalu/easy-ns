package container

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

func NewProcess(tty bool) (*exec.Cmd, *os.File) {
	readPipe, writePipi, err := NewPipe()
	if err != nil {
		logrus.Errorf("New pipe error %v", err)
		return nil, nil
	}

	self, err := os.Readlink("/proc/self/exe")
	if err != nil {
		logrus.Errorf("read link /proc/self/exe error %v", err)
		return nil, nil
	}

	logrus.Infof("read link %s", self)

	cmd := exec.Command(self, "init")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
		GidMappingsEnableSetgroups: true,
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.ExtraFiles = []*os.File{readPipe}
	cmd.Dir = "/root/busybox"
	return cmd, writePipi
}

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}
