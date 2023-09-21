package wrappers

import (
	"encoding/json"
	"os"
	"strings"
	"xs/pkg/io"
)

type BuildConfig struct {
	Dest string `json:"dest"`
}

func LoadNgConfig(confFile string) *BuildConfig {
	bc := &BuildConfig{}
	bytesFromFile, err := os.ReadFile(confFile)
	if err != nil {
		io.Panic(err)
	}
	err = json.Unmarshal([]byte(bytesFromFile), bc)
	if err != nil {
		io.Panic(err)
	}

	return bc
}

func LoadNgDest(dir string) string {
	confFile := dir + "/xs-package.json"
	config := LoadNgConfig(confFile)
	destPath := config.Dest
	destFixed := strings.Replace(destPath, "../../../../", "./", -1)
	return destFixed
}
