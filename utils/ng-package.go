package utils

import (
	"encoding/json"
	"strings"
)

type BuildConfig struct {
	Dest string `json:"dest"`
}

func LoadNgConfig(confFile string) *BuildConfig {
	bc := &BuildConfig{}
	bytesFromFile, err := ReadFile(confFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bytesFromFile), bc)
	if err != nil {
		panic(err)
	}

	return bc
}

func LoadNgDest(dir string) string {
	confFile := dir + "/ng-package.json"
	config := LoadNgConfig(confFile)
	destPath := config.Dest
	destFixed := strings.Replace(destPath, "../../../../", "./", -1)
	return destFixed
}
