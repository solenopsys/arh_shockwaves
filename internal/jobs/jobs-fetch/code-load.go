package jobs_fetch

import (
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/controllers"
	"xs/pkg/io"
)

type CodeLoad struct {
	cid         string
	packageName string
	targetDir   string
	sourceUrl   string
}

func (t *CodeLoad) Execute() *jobs.Result {
	source := controllers.NewSourceLoader()
	err := source.Load(t.cid, t.targetDir, t.sourceUrl)

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
		Description: "Code loaded " + t.packageName + " -> " + t.targetDir,
	}
}

func (t *CodeLoad) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: t.packageName + " -> " + t.targetDir,
		Short:       "Reddy",
	}

}

func NewCodeLoad(cid string, packageName string, targetDir string, sourceUrl string) *CodeLoad {
	return &CodeLoad{cid: cid, packageName: packageName, targetDir: targetDir, sourceUrl: sourceUrl}
}
