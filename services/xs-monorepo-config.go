package services

import (
	"encoding/json"
	"xs/utils"
)

/*
		{
		  "format": {
		    "name": "xs-backs",
		    "version": "1.0.0"
		  },
		  "modules": {
		    "platform": [
		      {
		        "directory": "invicta/helm-repository",
		        "git": "https://github.com/solenopsys/converged-microservice-helm-repository-invicta",
		        "load": true
		      },
	     {
	        "npm":"@solenopsys/lib-clusters",
	        "directory": "clusters",
	        "git": "https://github.com/solenopsys/sc-fl-helm",
	        "use": "git",
	        "load": true
	      },
	    {
	        "npm":"@solenopsys/lib-clusters",
	        "directory": "clusters",
	        "git": "https://github.com/solenopsys/sc-fl-helm",
	        "use": "npm",
	        "load": true
	      }
		    ]
		  }
		}
*/
type XsMonorepoFormat struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type XsMonorepoConfig struct {
	Format  XsMonorepoFormat              `json:"format"`
	Modules map[string][]XsMonorepoModule `json:"modules"`
}

type XsMonorepoModule struct {
	Name      string `json:"name"`
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

func (c *ConfLoader) LoadConfig() {
	config := &XsMonorepoConfig{}
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
	for section, group := range groups.Modules {
		for _, module := range group {
			println("Start load repository:", module.Name)
			path := section + "/" + c.targetDir + "/" + module.Directory
			utils.CloneGitRepository(module.Git, path)
		}
	}
}

const treeRepoLink = "https://github.com/solenopsys/treerepo-template.git"

func LoadBase() {
	println("Start load base\n")

	err := utils.CloneGitRepository(treeRepoLink, "./")
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
	loader.targetDir = "./"
	loader.SyncFunc = syncBack
	return &loader
}

func syncFront() {
	loader := NewFrontLoader()
	loader.LoadConfig()
	loader.SyncModules()
	InjectConfToJson(loader, "./front/tsconfig.base.json")
}

func syncBack() {
	backLoader := NewBackLoader()
	backLoader.LoadConfig()
	backLoader.SyncModules()
}
