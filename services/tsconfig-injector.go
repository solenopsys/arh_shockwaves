package services

import (
	"encoding/json"
	"os"
	"xs/utils"
)

func InjectConfToJson(c *ConfLoader, fileName string, filter string) {

	existingJSON, err := utils.ReadFile(fileName)
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

				println("Inject to package.json:", module.Npm, path)

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
