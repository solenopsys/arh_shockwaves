package configs

import (
	"encoding/json"
	"github.com/spf13/viper"
	"os"
	"xs/pkg/io"
)

type TsInjector struct {
	packageFileName string
	confData        map[string]any
}

func (i *TsInjector) Load() {
	existingJSON, err := os.ReadFile(i.packageFileName)
	if err != nil {
		io.Panic(err)
	}

	err = json.Unmarshal([]byte(existingJSON), &i.confData)
	if err != nil {
		io.Panic(err)
	}
}

func (i *TsInjector) Save() {
	newJSON, err := json.MarshalIndent(i.confData, "", "  ")
	if err != nil {
		io.Panic(err)
	}

	os.WriteFile(i.packageFileName, newJSON, 0644)
}

func NewTsInjector() *TsInjector {
	return &TsInjector{
		packageFileName: "." + viper.GetString("files.tsconfig"),
		confData:        make(map[string]any),
	}
}

func (i *TsInjector) loadPaths() map[string]any {
	var modulesConf map[string]interface{} = i.confData["compilerOptions"].(map[string]any)["paths"].(map[string]interface{})
	if modulesConf == nil {
		modulesConf = make(map[string]interface{})
	}
	return modulesConf
}

func (i *TsInjector) savePaths(paths map[string]any) {
	i.confData["compilerOptions"].(map[string]any)["paths"] = paths
}

func (i *TsInjector) InjectPackagesLinksTsconfigJson(packages map[string]string) {
	for modName, path := range packages {
		i.AddPackage(modName, path)
	}
}

func (i *TsInjector) AddPackage(npmName string, directory string) {
	modulesConf := i.loadPaths()
	tsFile := directory + "/src/index.ts"
	modulesConf[npmName] = []string{tsFile}

	i.savePaths(modulesConf)
}

func (i *TsInjector) RemovePackage(name string) {
	modulesConf := i.loadPaths()

	delete(modulesConf, name)

	i.savePaths(modulesConf)
}
