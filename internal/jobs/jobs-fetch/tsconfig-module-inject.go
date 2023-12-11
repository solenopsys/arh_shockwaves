package jobs_fetch

import (
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/tools"
)

type TsConfigModuleInject struct {
	packageName string
	targetDir   string
}

func (t *TsConfigModuleInject) Execute() *jobs.Result {
	injector := configs.NewTsInjector()
	res, rootDir := tools.FindWorkspaceRootDir()
	if !res {
		return &jobs.Result{
			Success:     false,
			Error:       nil,
			Description: "Workspace root dir not found",
		}
	}
	pt := tools.PathTools{}

	pt.SetBasePathPwd()
	pt.MoveTo(rootDir + "/converged")
	injector.Load()
	injector.AddPackage(t.packageName, t.targetDir)
	injector.Save()
	pt.MoveToBasePath()

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
		Key:         "tsconfig-module-inject-" + t.packageName,
	}

}

func NewTsConfigModuleInject(packageName string, targetDir string) jobs.PrintableJob {
	subDir := strings.Replace(targetDir, "converged/", "", 1) // todo move to const or change logic
	return &TsConfigModuleInject{packageName: packageName, targetDir: subDir}
}
