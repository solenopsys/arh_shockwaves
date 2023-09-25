package jobs_fetch

import (
	"xs/internal/jobs"
	"xs/pkg/controllers"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type TemplateLoad struct {
	packageName string
	targetDir   string
}

func (t *TemplateLoad) Execute() *jobs.Result {
	pinning := wrappers.NewPinning()
	repo, err := pinning.FindOne(t.packageName)
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}
	templateLoader := controllers.NewSourceLoader()
	err = templateLoader.Load(repo.Cid, t.targetDir, "")
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}
	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "Template loaded" + t.packageName + " to " + t.targetDir,
	}
}

func (t *TemplateLoad) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Template load " + t.packageName + " to " + t.targetDir,
		Short:       "Reddy",
	}
}

func NewTemplateLoad(packageName string, targetDir string) *TemplateLoad {
	return &TemplateLoad{packageName: packageName, targetDir: targetDir}
}
