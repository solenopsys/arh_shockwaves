package jobs_fetch

import (
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
)

type TsConfigModuleInject struct {
	packageName string
	targetDir   string
}

func (t *TsConfigModuleInject) Execute() *jobs.Result {
	injector := configs.NewTsInjector()
	injector.Load()
	injector.AddPackage(t.packageName, t.targetDir)
	injector.Save()

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Link injected in tsconfig " + t.packageName + " -> " + t.targetDir,
	}
}

func (t *TsConfigModuleInject) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Inject to tsconfig link for:  " + t.packageName,
		Name:        t.packageName,
	}

}

func NewTsConfigModuleInject(packageName string, targetDir string) jobs.PrintableJob {
	subDir := strings.Replace(targetDir, "frontends/", "", 1) // todo move to const or change logic
	return &TsConfigModuleInject{packageName: packageName, targetDir: subDir}
}
