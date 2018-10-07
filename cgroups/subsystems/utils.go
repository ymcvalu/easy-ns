package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

func FindCgronpMountPoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return strings.Split(fields[4], ",")[0]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		logrus.Errorf("when find mount point of subsystemL %s", err.Error())
		return ""
	}
	return ""
}

func GetCgroupPath(sbusystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgronpMountPoint(sbusystem)
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (os.IsNotExist(err) && autoCreate) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err != nil {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("cgroup path error: %v", err)
	}
}
