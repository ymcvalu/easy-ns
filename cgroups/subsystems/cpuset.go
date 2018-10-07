package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpusetSubSystem struct {
}

//type assert
var _ Subsystem = (*CpusetSubSystem)(nil)

func (cs *CpusetSubSystem) Name() string {
	return "cpuset"
}

func (cs *CpusetSubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(cs.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail: %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	}
}

func (cs *CpusetSubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(cs.Name(), cgroupPath, false); err == nil {
		return os.RemoveAll(subsysCgroupPath)
	} else {
		return err
	}
}

func (cs *CpusetSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(cs.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cpuset.cpus"), []byte(res.CPUShare), 0644); err != nil {
				return fmt.Errorf("set cgroup cpuset fail: %v", err)
			}

		}
		return nil
	} else {
		return err
	}
}
