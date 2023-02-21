package dev

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

type ConfLoader struct {
	configName string
}

func (c *ConfLoader) loadConfig() *[]ModuleGroup {

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

func (c *ConfLoader) syncModules() {
	groups := *c.loadConfig()
	for _, group := range groups {
		for _, module := range group.Modules {
			println("Load repository: ", module.Name)
			path := "./front/packages/" + group.Dir + "/" + module.Directory
			utils.CloneGitRepository(module.Git, path)
		}
	}
}

func NewLoader() *ConfLoader {
	loader := ConfLoader{}
	loader.configName = "./config/front-modules.json"
	return &loader
}

func (c *ConfLoader) injectConfiguration() {
	// c.loadModules()
}
