package services

import (
	"encoding/json"
	"os"
	"xs/utils"
)

func InjectConfToJson(c *ConfLoader, fileName string) {

	existingJSON, err := utils.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var confData map[string]any
	err = json.Unmarshal([]byte(existingJSON), &confData)
	if err != nil {
		panic(err)
	}

	modulesConf := make(map[string][]string)

	groups := *c.data

	for section, group := range groups.Groups {
		for _, module := range group {
			path := c.targetDir + "/" + section + "/" + module.Directory
			tsFile := path + "/src/index.ts"
			println("Inject to config:", module.Npm, tsFile)
			modulesConf[module.Npm] = []string{tsFile}
		}
	}

	confData["compilerOptions"].(map[string]any)["paths"] = modulesConf

	newJSON, err := json.MarshalIndent(confData, "", "  ")
	if err != nil {
		panic(err)
	}

	os.WriteFile(fileName, newJSON, 0644)
}
