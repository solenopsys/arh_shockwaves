package configs

import (
	"encoding/json"
	"sync"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/wrappers"
)

const FRONT = "front"
const BACK = "back"

var FileTypeMapping = map[string]string{
	"xs-fronts": FRONT,
	"xs-backs":  BACK,
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
	Tags      []string `json:"tags"`
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
	fileData, err := tools.ReadFile(fileName)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), config)
	} else {
		io.Fatal(err)
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
			wrappers.CloneGitRepository(module.Git, path, true, false)
		}
	}
}

func LoadWorkspace(monorepoLink string) {

	path := "."
	wg := sync.WaitGroup{}
	wg.Add(1)
	err := wrappers.CloneGitRepository(monorepoLink, path, false, false)
	if err != nil {
		io.Panic(err)
	} else {
		gitDir := path + "/.git"
		err := tools.DeleteDir(gitDir)
		if err != nil {
			io.Panic(err)
		}
	}
}

func LoadBase(monorepoLink string) {
	io.Println("Tags base\n")

	err := wrappers.CloneGitRepository(monorepoLink, ".", false, false)
	if err != nil {
		io.Panic(err)
	} else {
		io.Println("Base loaded\n")
	}
}

func NewFrontLoader(path string, tsConfig bool) *ConfLoader {
	loader := ConfLoader{}
	loader.configName = path + "/xs-treerepo.json" // todo move to const
	loader.targetDir = path
	loader.LoadConfig()
	loader.SyncFunc = func() {
		loader.SyncModules()
		if tsConfig {
			InjectConfToTsconfigJson(&loader, "./tsconfig.develop.json")
		} else {
			InjectToPackageJson(&loader, "./package.json", "packages")
		}
	}
	return &loader
}

func NewBackLoader(path string) *ConfLoader {
	loader := ConfLoader{}
	loader.configName = path + "/xs-treerepo.json" // todo move to const
	loader.targetDir = path
	loader.LoadConfig()
	loader.SyncFunc = func() {
		loader.SyncModules()
	}
	return &loader
}
