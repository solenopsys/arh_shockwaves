package configs

import (
	"encoding/json"
	"os"
	"xs/pkg/io"
)

func InjectPackagesLinksTsconfigJson(packages map[string]string, packageJsonFileName string) {
	existingJSON, err := os.ReadFile(packageJsonFileName)
	if err != nil {
		io.Panic(err)
	}
	var confData map[string]any
	err = json.Unmarshal([]byte(existingJSON), &confData)
	if err != nil {
		io.Panic(err)
	}

	compillerOptions := confData["compilerOptions"].(map[string]any)
	var modulesConf map[string]interface{} = compillerOptions["paths"].(map[string]interface{})
	if modulesConf == nil {
		modulesConf = make(map[string]interface{})
	}

	for modName, path := range packages {
		tsFile := path + "/src/index.ts"

		npm := modName
		io.Println("Inject to config:", npm, tsFile)
		modulesConf[npm] = []string{tsFile}
	}

	compillerOptions["paths"] = modulesConf

	newJSON, err := json.MarshalIndent(confData, "", "  ")
	if err != nil {
		io.Panic(err)
	}

	os.WriteFile(packageJsonFileName, newJSON, 0644)
}
