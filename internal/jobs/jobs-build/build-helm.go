package jobs_build

import (
	"path/filepath"
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type BuildHelm struct {
	params map[string]string
}

func (b *BuildHelm) Execute() *jobs.Result {

	path := b.params["path"]
	parentDir := "helm"
	pathHelmDir := path + "/" + parentDir
	absolutPath, _ := filepath.Abs(pathHelmDir)
	io.Println("path", absolutPath)

	arch := wrappers.ArchiveDir(absolutPath, parentDir)

	io.Println("archive size", len(arch))
	wrappers.PushDir(arch) // todo extract to push job or step inside

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "BuildHelm executed",
	}
}

func (b *BuildHelm) Description() string {
	return "Build helm " + b.params["path"]
}

func NewBuildHelm(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &BuildHelm{
		params: params,
	}
}
