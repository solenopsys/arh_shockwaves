package jobs_fetch

import (
	"os"
	"xs/internal/configs"
	"xs/internal/jobs"
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

func (t *CodeRemove) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Remove " + t.packageName + " from " + t.targetDir,
		Name:        t.packageName,
		Key:         "code-remove-" + t.packageName,
	}

}

func NewCodeRemove(packageName string, targetDir string) *CodeRemove {
	return &CodeRemove{packageName: packageName, targetDir: targetDir}
}
