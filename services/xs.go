package services

import (
	"errors"
	"github.com/tendermint/tendermint/libs/os"
)

type XsManager struct {
	config *XsMonorepoConfig
}

func (x *XsManager) load(fileName string) error {
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

func (x *XsManager) extract(group string, name string) *XsMonorepoModule {
	//println("Scan modules")
	groups := x.config.Groups
	modules := groups[group]

	for _, module := range modules {
		//println("Module name", module.Name)
		if module.Name == name || module.Npm == name {
			return module
		}
	}
	return nil
}

func (x *XsManager) extractGroup(group string) []*XsMonorepoModule {
	groups := x.config.Groups
	return groups[group]
}

func ExtractModule(m string, groupDir string, rType string) (*XsMonorepoModule, error) {
	var okModule *XsMonorepoModule
	var err error
	fileName := "./xs.json"
	xm := &XsManager{}
	err = xm.load(fileName)
	if err != nil {
		err = errors.New("xs.json not found, directory not initialized")
	} else {
		config := LoadConfigFile(fileName)
		repoType := FileTypeMapping[config.Format.Name]

		if repoType != rType {
			err = errors.New("Not applicable for " + rType + " monorepo")
		} else if repoType == rType {
			okModule = xm.extract(groupDir, m)

			if okModule == nil {
				err = errors.New("Module not found")
			} else {
				println("Ok module found")
			}
		} else {
			err = errors.New("Invalid xs.json, config type only xs-fronts or xs-backs allowed")
		}

	}

	return okModule, err

}
