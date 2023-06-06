package configs

import (
	"encoding/json"
	"os"
	"xs/pkg/io"
	"xs/pkg/tools"
)

func InjectToPackageJson(c *ConfLoader, fileName string, filter string) {

	existingJSON, err := tools.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var confData map[string]any
	err = json.Unmarshal([]byte(existingJSON), &confData)
	if err != nil {
		panic(err)
	}

	groups := *c.data

	for section, group := range groups.Groups {
		if section == filter {
			for _, module := range group {
				path := "file:" + c.targetDir + "/" + section + "/" + module.Directory + "/dist"

				io.Println("Inject to package.json:", module.Npm, path)

				confData["dependencies"].(map[string]any)[module.Npm] = path
			}
		}
	}

	newJSON, err := json.MarshalIndent(confData, "", "  ")
	if err != nil {
		panic(err)
	}

	os.WriteFile(fileName, newJSON, 0644)
}

type Packages struct {
}

func InjectConfToTsconfigJson(c *ConfLoader, fileName string) {

	existingJSON, err := tools.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var confData map[string]any
	err = json.Unmarshal([]byte(existingJSON), &confData)
	if err != nil {
		panic(err)
	}

	modulesConf := make(map[string][]string)

	groups := c.data.Groups
	for key, group := range groups {
		for _, module := range group {

			path := c.targetDir + "/" + key + "/" + module.Directory
			tsFile := path + "/src/public_api.ts"

			npm := module.Npm
			io.Println("Inject to config:", npm, tsFile)
			modulesConf[npm] = []string{tsFile}
		}
	}

	confData["compilerOptions"].(map[string]any)["paths"] = modulesConf

	newJSON, err := json.MarshalIndent(confData, "", "  ")
	if err != nil {
		panic(err)
	}

	os.WriteFile(fileName, newJSON, 0644)
}
