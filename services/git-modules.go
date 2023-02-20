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
	repoType   string
	configName string
	targetDir  string
	data       *[]GitModuleGroup
}

const treeRepoLink = "https://github.com/solenopsys/treerepo-template.git"

func LoadBase() {
	println("Start load base\n")

	err := utils.CloneGitRepository(treeRepoLink, "./")
	if err != nil {
		println("Error: %", err)
		panic(err)
	}

}

func (c *ConfLoader) LoadConfig() {
	config := &[]GitModuleGroup{}
	fileName := c.configName
	fileData, err := utils.ReadFile(fileName)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), config)
	}
	if err != nil {
		panic(err)
	}
	c.data = config
}

func (c *ConfLoader) SyncModules() {
	groups := *c.data
	for _, group := range groups {
		for _, module := range group.Modules {
			println("Start load repository:", module.Name)
			path := c.targetDir + "/" + group.Dir + "/" + module.Directory
			utils.CloneGitRepository(module.Git, path)
		}
	}
}

func NewFrontLoader() *ConfLoader {
	loader := ConfLoader{}
	loader.configName = "./config/front-modules.json"
	loader.targetDir = "./front/packages"
	return &loader
}

func NewBackLoader() *ConfLoader {
	loader := ConfLoader{}
	loader.configName = "./config/back-modules.json"
	loader.targetDir = "./back"
	return &loader
}

func (c *ConfLoader) injectConfiguration() {
	// c.loadModules()
}
func syncFront() {
	loader := NewFrontLoader()
	loader.LoadConfig()
	// loader.SyncModules()
	InjectConfToJson(loader, "./front/tsconfig.base.json")
}

func syncBack() {
	backLoader := NewBackLoader()
	backLoader.LoadConfig()
	backLoader.SyncModules()
}

func SyncAllModules() {
	syncFront()
	syncBack()
}
