package configs

import (
	"encoding/json"
	"xs/pkg/io"
	"xs/pkg/tools"
)

type Module struct {
	Name      string `json:"name"`
	Directory string `json:"directory"`
	Git       string `json:"git"`
}

type ModuleGroup struct {
	Dir     string   `json:"dir"`
	Modules []Module `json:"modules"`
}

type XsConfLoader struct {
	configName string
}

func (c *XsConfLoader) loadConfig() *[]ModuleGroup {

	config := &[]ModuleGroup{}
	fileName := c.configName
	fileData, err := tools.ReadFile(fileName)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), config)
	}
	if err != nil {
		io.Panic(err)
	}

	return config
}

func NewLoader() *XsConfLoader {
	loader := XsConfLoader{}
	loader.configName = "./config/front-modules.json"
	return &loader
}

func (c *XsConfLoader) injectConfiguration() {
	// c.loadModules()
}
