package services

import (
	"encoding/json"
	"sync"
	"xs/utils"
)

var FileTypeMapping = map[string]string{
	"xs-fronts": "front",
	"xs-backs":  "back",
}

type XsMonorepoFormat struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type XsMonorepoConfig struct {
	Format XsMonorepoFormat               `json:"format"`
	Groups map[string][]*XsMonorepoModule `json:"groups"`
}

type XsMonorepoModule struct {
	Directory string   `json:"directory"`
	Git       string   `json:"repository"`
	Load      []string `json:"scopes"`
	Npm       string   `json:"package"`
	Name      string   `json:"name"`
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

func LoadWorkspace(monorepoLink string) {

	path := "."
	wg := sync.WaitGroup{}
	wg.Add(1)
	err := utils.CloneGitRepository(monorepoLink, path, false, false)
	if err != nil {
		panic(err)
	} else {
		gitDir := path + "/.git"
		err := utils.DeleteDir(gitDir)
		if err != nil {
			panic(err)
		}
	}
}

func LoadBase(monorepoLink string) {
	println("Load base\n")

	err := utils.CloneGitRepository(monorepoLink, ".", false, false)
	if err != nil {
		panic(err)
	}
}

func NewFrontLoader(path string) *ConfLoader {
	loader := ConfLoader{}
	loader.configName = path + "/xs.json"
	loader.targetDir = path
	loader.LoadConfig()
	loader.SyncFunc = func() {
		loader.SyncModules()
		//InjectToPackageJson(&loader, "./package.json", "libraries")
		InjectConfToTsconfigJson(&loader, "./tsconfig.develop.json")
	}
	return &loader
}

func NewBackLoader(path string) *ConfLoader {
	loader := ConfLoader{}
	loader.configName = path + "/xs.json"
	loader.targetDir = path
	loader.LoadConfig()
	loader.SyncFunc = func() {
		loader.SyncModules()
	}
	return &loader
}
