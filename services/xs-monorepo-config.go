package services

import (
	"encoding/json"
	"xs/utils"
)

type XsMonorepoFormat struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type XsMonorepoConfig struct {
	Format XsMonorepoFormat              `json:"format"`
	Groups map[string][]XsMonorepoModule `json:"groups"`
}

type XsMonorepoModule struct {
	Directory string `json:"directory"`
	Git       string `json:"git"`
	Load      bool   `json:"load"`
	Npm       string `json:"npm"`
	Use       string `json:"use"`
}

type ConfLoader struct {
	repoType   string
	configName string
	targetDir  string
	data       *XsMonorepoConfig
	SyncFunc   func()
}

func LoadConfigFile(fileName string) *XsMonorepoConfig {
	config := &XsMonorepoConfig{}
	fileData, err := utils.ReadFile(fileName)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), config)
	}
	if err != nil {
		panic(err)
	}
	return config
}

func (c *ConfLoader) LoadConfig() {
	c.data = LoadConfigFile(c.configName)
}

func (c *ConfLoader) SyncModules() {
	groups := *c.data
	for section, group := range groups.Groups {
		for _, module := range group {
			path := c.targetDir + "/" + section + "/" + module.Directory
			utils.CloneGitRepository(module.Git, path, true, false)
		}
	}
}

func LoadBase(monorepoLink string) {
	println("Load base\n")

	err := utils.CloneGitRepository(monorepoLink, "./", false, false)
	if err != nil {
		panic(err)
	}

}

func NewFrontLoader() *ConfLoader {
	loader := ConfLoader{}
	loader.configName = "./xs.json"
	loader.targetDir = "./packages"
	loader.SyncFunc = syncFront
	return &loader
}

func NewBackLoader() *ConfLoader {
	loader := ConfLoader{}
	loader.configName = "./xs.json"
	loader.targetDir = ""
	loader.SyncFunc = syncBack
	return &loader
}

func syncFront() {
	loader := NewFrontLoader()
	loader.LoadConfig()
	loader.SyncModules()
	InjectConfToJson(loader, "./tsconfig.base.json")
}

func syncBack() {
	backLoader := NewBackLoader()
	backLoader.LoadConfig()
	backLoader.SyncModules()
}
