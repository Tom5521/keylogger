package autostart_test

import (
	"testing"

	"github.com/Tom5521/keylogger/internal/autostart"
)

func TestMakeShortcut(t *testing.T) {
	app := autostart.Make()
	if autostart.Check(app) {
		err := app.Disable()
		if err != nil {
			t.Fail()
		}
	} else {
		err := app.Enable()
		if err != nil {
			t.Fail()
		}
	}
}
