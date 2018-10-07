package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSubSystem struct{}

//type assert
var _ Subsystem = (*CpuSubSystem)(nil)

func (cpu *CpuSubSystem) Name() string {
	return "cpu"
}

func (cpu *CpuSubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(cpu.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail: %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	}
}

func (cpu *CpuSubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(cpu.Name(), cgroupPath, false); err == nil {
		return os.RemoveAll(subsysCgroupPath)
	} else {
		return err
	}
}

func (cpu *CpuSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(cpu.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cpu.shares"), []byte(res.CPUShare), 0644); err != nil {
				return fmt.Errorf("set cgroup cpu fail: %v", err)
			}

		}
		return nil
	} else {
		return err
	}
}
