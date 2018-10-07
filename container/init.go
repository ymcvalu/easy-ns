package container

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
)

func RunContainerInitProcess() error {
	cmds := readUserCommand()
	if cmds == nil || len(cmds) == 0 {
		return fmt.Errorf("Run container get user command error,command is empty")
	}

	setUpMount()

	path, err := exec.LookPath(cmds[0])
	if err != nil {
		logrus.Errorf("Exec look path error %v", err)
		return err
	}
	logrus.Infof("Find path %s", path)
	if err := syscall.Exec(path, cmds[0:], os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}

func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		logrus.Errorf("init read pipe error %v", err)
		return nil
	}
	msgStr := string(msg)
	return strings.Split(msgStr, " ")
}

func setUpMount() {
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Errorf("get current work dir error: %v", err)
		return
	}

	if err := mountProc(pwd); err != nil {
		logrus.Errorf("mount proc error %v", err)
		return
	}

	if err := mountDev(pwd); err != nil {
		logrus.Errorf("mount dev error %v", err)
		return
	}

	if err := pivotRoot(pwd); err != nil {
		logrus.Errorf("pivot root error %v", err)
		return
	}

}

func pivotRoot(root string) error {

	putold := filepath.Join(root, "/.pivot_root")

	//create a mount point for new root
	if err := syscall.Mount(root, root, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("Mount rootfs to itself error: %v", err)
	}

	if err := os.Mkdir(putold, 0700); err != nil && !os.IsExist(err) {
		return err
	}

	if err := syscall.PivotRoot(root, putold); err != nil {
		return fmt.Errorf("pivot root %v", err)
	}

	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir / %v", err)
	}

	putold = "/.pivot_root"
	if err := syscall.Unmount(putold, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount old root %v", err)
	}

	return os.Remove(putold)
}

func mountProc(root string) error {
	path := filepath.Join(root, "/proc")
	return syscall.Mount("proc", path, "proc", uintptr(0), "")
}

func mountDev(root string) error {
	path := filepath.Join(root, "/dev")
	return syscall.Mount("tmpfs", path, "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
}
