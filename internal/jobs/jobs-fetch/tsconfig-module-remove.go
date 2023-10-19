package jobs_fetch

import (
	"xs/internal/configs"
	"xs/internal/jobs"
)

type TsConfigModuleRemove struct {
	packageName string
}

func (t *TsConfigModuleRemove) Execute() *jobs.Result {
	injector := configs.NewTsInjector()
	injector.Load()
	injector.RemovePackage(t.packageName)
	injector.Save()

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Link removed from tsconfig for package:  " + t.packageName,
	}
}

func (t *TsConfigModuleRemove) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Remove link from tsconfig:  " + t.packageName,
		Name:        t.packageName,
		Key:         "tsconfig-module-remove-" + t.packageName,
	}

}

func NewTsConfigModuleRemove(packageName string, targetDir string) jobs.PrintableJob {
	return &TsConfigModuleRemove{packageName: packageName}
}
