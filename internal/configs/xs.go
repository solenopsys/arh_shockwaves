package configs

import (
	"errors"
	"fmt"
	"github.com/tendermint/tendermint/libs/os"
	"regexp"
	"strings"
	"xs/pkg/io"
)

type XsManager struct {
	config *XsMonorepoConfig
}

func (x *XsManager) Load(fileName string) error {
	var err error
	exists := os.FileExists(fileName)
	if exists {
		x.config = LoadConfigFile(fileName)
	} else {
		var name = "xs-treerepo.json" // todo move to global const
		err = errors.New(name + " not found, directory not initialized")
	}
	return err
}

func (x *XsManager) ExtractModule(group string, name string) *XsMonorepoModule {
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

func (x *XsManager) FilterLibs(filter string, group string) []*XsMonorepoModule {
	groups := x.ExtractGroup(group)
	var filtered []*XsMonorepoModule = []*XsMonorepoModule{}
	for _, module := range groups {
		name := module.Name
		pattern := strings.Replace(filter, "*", ".*", -1)
		matched, err := regexp.MatchString(pattern, name)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if matched {
			filtered = append(filtered, module)
		}
	}

	io.Println("Found  lib count:", len(filtered))
	return filtered
}

func ExtractModule(m string, groupDir string, rType string) (*XsMonorepoModule, error) {
	var okModule *XsMonorepoModule
	var err error
	fileName := "./xs-treerepo.json" //todo move to const
	xm := &XsManager{}
	err = xm.Load(fileName)
	if err != nil {
		err = errors.New(fileName + " not found, directory not initialized")
	} else {
		config := LoadConfigFile(fileName)
		repoType := FileTypeMapping[config.Format.Name]

		if repoType != rType {
			err = errors.New("Not applicable for " + rType + " monorepo")
		} else if repoType == rType {
			okModule = xm.ExtractModule(groupDir, m)

			if okModule == nil {
				err = errors.New("Module not found")
			} else {
				io.Println("Ok module found")
			}
		} else {
			err = errors.New("Invalid " + fileName + ", config type only xs-fronts or xs-backs allowed")
		}

	}

	return okModule, err

}
