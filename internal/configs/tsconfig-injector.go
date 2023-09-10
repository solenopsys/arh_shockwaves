package configs

import (
	"encoding/json"
	"os"
	"xs/pkg/io"
	"xs/pkg/tools"
)

func InjectConfToTsconfigJson(packages map[string]string, packageJsonFileName string) {
	existingJSON, err := tools.ReadFile(packageJsonFileName)
	if err != nil {
		io.Panic(err)
	}
	var confData map[string]any
	err = json.Unmarshal([]byte(existingJSON), &confData)
	if err != nil {
		io.Panic(err)
	}

	modulesConf := make(map[string][]string)

	for modName, path := range packages {
		tsFile := path + "/src/index.ts"

		npm := modName
		io.Println("Inject to config:", npm, tsFile)
		modulesConf[npm] = []string{tsFile}
	}

	confData["compilerOptions"].(map[string]any)["paths"] = modulesConf

	newJSON, err := json.MarshalIndent(confData, "", "  ")
	if err != nil {
		io.Panic(err)
	}

	os.WriteFile(packageJsonFileName, newJSON, 0644)
}
