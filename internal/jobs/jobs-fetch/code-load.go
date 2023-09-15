package jobs_fetch

import (
	"xs/internal/configs"
	"xs/internal/jobs"
)

type CodeLoad struct {
	cid         string
	packageName string
	targetDir   string
}

func (t *CodeLoad) Execute() *jobs.Result {
	source := configs.NewSourceLoader()
	err := source.Load(t.cid, t.targetDir)

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	ws, err := configs.GetInstanceWsManager()

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	ws.AddModule(t.packageName, t.targetDir)

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Code loaded" + t.packageName + " to " + t.targetDir,
	}
}

func (t *CodeLoad) Description() string {
	return t.packageName + " -> " + t.targetDir
}

func NewCodeLoad(cid string, packageName string, targetDir string) *CodeLoad {
	return &CodeLoad{cid: cid, packageName: packageName, targetDir: targetDir}
}
