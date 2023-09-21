package configs

import (
	"encoding/json"
	"os"
	"xs/pkg/io"
)

func InjectToPackageJson(packages map[string]string, fileName string) {
	existingJSON, err := os.ReadFile(fileName)
	if err != nil {
		io.Panic(err)
	}
	var confData map[string]any
	err = json.Unmarshal([]byte(existingJSON), &confData)
	if err != nil {
		io.Panic(err)
	}

	for modName, path := range packages {
		io.Println("Inject to package.json:", modName, path)

		confData["dependencies"].(map[string]any)[modName] = path
	}

	newJSON, err := json.MarshalIndent(confData, "", "  ")
	if err != nil {
		io.Panic(err)
	}

	os.WriteFile(fileName, newJSON, 0644)
}
