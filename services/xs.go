package services

import (
	"errors"
	"github.com/tendermint/tendermint/libs/os"
)

func ExtractModule(m string, groupDir string) (XsMonorepoModule, error) {
	var okModule XsMonorepoModule
	var err error
	fileName := "./xs.json"
	exists := os.FileExists(fileName)
	if exists {
		config := LoadConfigFile(fileName)
		repoType := FileTypeMapping[config.Format.Name]

		if repoType == "front" {
			err = errors.New("Not applicable for front monorepo")
		} else if repoType == "back" {

			groups := config.Groups
			modules := groups[groupDir]

			var ok = false
			for _, module := range modules {
				if module.Directory == m {
					okModule = module
					ok = true
				}
			}

			if ok {
				println("Ok module found")

			} else {
				err = errors.New("Module not found")
			}
		} else {
			err = errors.New("Invalid xs.json, config type only xs-fronts or xs-backs allowed")
		}

	} else {
		err = errors.New("xs.json not found, directory not initialized")
	}

	return okModule, err

}
