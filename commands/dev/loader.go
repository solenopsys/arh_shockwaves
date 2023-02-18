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
	file := "front-modules.json"
	config := &[]ModuleGroup{}
	err := json.Unmarshal([]byte(file), config)
	if err != nil {
		panic(err)
	}
	return config
}

func (c *ConfLoader) loadBase() {
	utils.CloneGitRepository("https://github.com/solenopsys/treerepo-template", "./")
}

func (c *ConfLoader) syncModules() {
	groups := *c.loadConfig()
	for _, group := range groups {
		for _, module := range group.Modules {
			println("Start load: %", module.Name)
			go utils.CloneGitRepository(module.Git, "./packages/"+group.Dir+"/"+module.Directory)
		}
	}
}

func NewLoader() *ConfLoader {
	loader := ConfLoader{}
	loader.configName = "config/front-modules.json"
	return &loader
}

func (c *ConfLoader) injectConfiguration() {
	// c.loadModules()
}
