package wrappers

import (
	"encoding/json"
	"os"
	"strings"
	"xs/pkg/io"
)

type Config struct {
	Projects map[string]Project `json:"projects"`
}

type Project struct {
	Architect Architect `json:"architect"`
}

type Architect struct {
	Build Build `json:"build"`
}

type Build struct {
	Options Options `json:"options"`
}

type Options struct {
	OutputPath string `json:"outputPath"`
}

func LoadAngConfig(confFile string) *Config {
	conf := &Config{}
	bytesFromFile, err := os.ReadFile(confFile)
	if err != nil {
		io.Panic(err)
	}
	err = json.Unmarshal([]byte(bytesFromFile), conf)
	if err != nil {
		io.Panic(err)
	}

	return conf
}

func LoadAngularDest(projectName string, dir string) string {
	confFile := dir + "/angular.json"
	config := LoadAngConfig(confFile)
	destPath := config.Projects[projectName].Architect.Build.Options.OutputPath
	destFixed := strings.Replace(destPath, "../../../", "./", -1)
	return destFixed
}
