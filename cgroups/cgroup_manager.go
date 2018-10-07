package cgroups

import (
	"easy-ns/cgroups/subsystems"

	"github.com/sirupsen/logrus"
)

type CgroupManager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

func (m *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range subsystems.SubsystemIns {
		subSysIns.Set(m.Path, res)
	}
	return nil
}

func (m *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subsystems.SubsystemIns {
		subSysIns.Apply(m.Path, pid)
	}
	return nil
}

func (m *CgroupManager) Destroy() error {
	for _, subSysIns := range subsystems.SubsystemIns {
		if err := subSysIns.Remove(m.Path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}
