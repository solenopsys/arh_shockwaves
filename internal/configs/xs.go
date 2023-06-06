package configs

import (
	"errors"
	"github.com/tendermint/tendermint/libs/os"
	"xs/pkg/io"
)

type XsManager struct {
	config *XsMonorepoConfig
}

func (x *XsManager) Load(fileName string) error {
	var err error
	//fileName := "./xs.json"
	exists := os.FileExists(fileName)
	if exists {
		x.config = LoadConfigFile(fileName)
	} else {
		err = errors.New("xs.json not found, directory not initialized")
	}
	return err
}

func (x *XsManager) Extract(group string, name string) *XsMonorepoModule {
	//io.Println("Scan modules")
	groups := x.config.Groups
	modules := groups[group]

	for _, module := range modules {
		//io.Println("Module name", module.Name)
		if module.Name == name || module.Npm == name {
			return module
		}
	}
	return nil
}

func (x *XsManager) ExtractGroup(group string) []*XsMonorepoModule {
	groups := x.config.Groups
	return groups[group]
}

func ExtractModule(m string, groupDir string, rType string) (*XsMonorepoModule, error) {
	var okModule *XsMonorepoModule
	var err error
	fileName := "./xs.json"
	xm := &XsManager{}
	err = xm.Load(fileName)
	if err != nil {
		err = errors.New("xs.json not found, directory not initialized")
	} else {
		config := LoadConfigFile(fileName)
		repoType := FileTypeMapping[config.Format.Name]

		if repoType != rType {
			err = errors.New("Not applicable for " + rType + " monorepo")
		} else if repoType == rType {
			okModule = xm.Extract(groupDir, m)

			if okModule == nil {
				err = errors.New("Module not found")
			} else {
				io.Println("Ok module found")
			}
		} else {
			err = errors.New("Invalid xs.json, config type only xs-fronts or xs-backs allowed")
		}

	}

	return okModule, err

}
