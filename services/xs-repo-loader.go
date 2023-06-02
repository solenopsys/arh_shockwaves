package services

import (
	"encoding/json"
	"xs/utils"
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
	fileData, err := utils.ReadFile(fileName)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), config)
	}
	if err != nil {
		panic(err)
	}

	return config
}

func (c *XsConfLoader) syncModules() {
	groups := *c.loadConfig()
	for _, group := range groups {
		for _, module := range group.Modules {
			println("Load repository: ", module.Name)
			path := "./front/packages/" + group.Dir + "/" + module.Directory
			utils.CloneGitRepository(module.Git, path, true, false)
		}
	}
}

func NewLoader() *XsConfLoader {
	loader := XsConfLoader{}
	loader.configName = "./config/front-modules.json"
	return &loader
}

func (c *XsConfLoader) injectConfiguration() {
	// c.loadModules()
}
