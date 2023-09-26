package jobs_fetch

import (
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/io"
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

func (t *TsConfigModuleRemove) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Remove link from tsconfig:  " + t.packageName,
		Short:       "Reddy",
	}

}

func NewTsConfigModuleRemove(packageName string, targetDir string) jobs.PrintableJob {
	return &TsConfigModuleInject{packageName: packageName}
}
