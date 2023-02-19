package services

import (
	"encoding/json"
	"xs/utils"
)

type GitModule struct {
	Name      string `json:"name"`
	Directory string `json:"directory"`
	Git       string `json:"git"`
}

type GitModuleGroup struct {
	Dir     string      `json:"dir"`
	Modules []GitModule `json:"modules"`
}

type ConfLoader struct {
	configName string
}

func (c *ConfLoader) loadConfig() *[]GitModuleGroup {

	config := &[]GitModuleGroup{}
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

func (c *ConfLoader) LoadBase() {
	println("Start load base\n")

	err := utils.CloneGitRepository("https://github.com/solenopsys/treerepo-template.git", "./")
	if err != nil {
		println("Error: %", err)
		panic(err)
	}

}

func (c *ConfLoader) SyncModules() {
	groups := *c.loadConfig()
	for _, group := range groups {
		for _, module := range group.Modules {
			println("Start load repository: ", module.Name)
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
