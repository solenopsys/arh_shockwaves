package jobs_fetch

import (
	"xs/internal/jobs"
	"xs/internal/services"
	"xs/pkg/controllers"
	"xs/pkg/io"
)

type TemplateLoad struct {
	packageName string
	targetDir   string
}

func (t *TemplateLoad) Execute() *jobs.Result {
	pinning := services.NewPinningRequests()
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
