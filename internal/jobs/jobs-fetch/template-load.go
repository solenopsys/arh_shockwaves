package jobs_fetch

import (
	"xs/internal/configs"
	"xs/internal/jobs"
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
			Err:     err,
		}
	}
	tl := configs.NewSourceLoader()
	tl.Load(repo.Cid, t.targetDir)
	return &jobs.Result{
		Success:     true,
		Err:         nil,
		Description: "Template loaded" + t.packageName + " to " + t.targetDir,
	}
}

func (t *TemplateLoad) Description() string {
	return "Template load " + t.packageName + " to " + t.targetDir
}

func NewTemplateLoad(packageName string, targetDir string) *TemplateLoad {
	return &TemplateLoad{packageName: packageName, targetDir: targetDir}
}
