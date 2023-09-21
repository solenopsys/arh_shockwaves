package jobs_fetch

import (
	"xs/internal/configs"
	"xs/internal/jobs"
)

type TsConfigModuleInject struct {
	packageName string
	targetDir   string
}

func (t *TsConfigModuleInject) Execute() *jobs.Result {

	packages := map[string]string{
		t.packageName: t.targetDir,
	}
	configs.InjectPackagesLinksTsconfigJson(packages, "."+configs.GetInstanceConfManager().Conf.Files.TsConfig)

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Code loaded" + t.packageName + " to " + t.targetDir,
	}
}

func (t *TsConfigModuleInject) Description() string {
	return "inject to tsconfig link for:  " + t.packageName
}

func NewTsConfigModuleInject(packageName string, targetDir string) jobs.PrintableJob {
	return &TsConfigModuleInject{packageName: packageName, targetDir: targetDir}
}
