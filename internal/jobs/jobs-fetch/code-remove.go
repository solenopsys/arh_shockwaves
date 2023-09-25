package jobs_fetch

import (
	"os"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/io"
)

type CodeRemove struct {
	packageName string
	targetDir   string
}

func (t *CodeRemove) Execute() *jobs.Result {
	err := os.RemoveAll(t.targetDir)

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	confManager, err := configs.GetInstanceWsManager()
	confManager.RemoveModule(t.packageName)

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Code removed " + t.packageName + " -> " + t.targetDir,
	}
}

func (t *CodeRemove) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Remove " + t.packageName + " from " + t.targetDir,
		Short:       "Reddy",
	}

}

func NewCodeRemove(packageName string, targetDir string) *CodeRemove {
	return &CodeRemove{packageName: packageName, targetDir: targetDir}
}
