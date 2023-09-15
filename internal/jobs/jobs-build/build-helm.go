package jobs_build

import (
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type BuildHelm struct {
	params map[string]string
}

func (b *BuildHelm) Execute() *jobs.Result {
	name := b.params["name"]
	path := b.params["path"]

	io.Println("path", path)
	arch := wrappers.ArchiveDir(path, name)

	io.Println("archive size", len(arch))
	wrappers.PushDir(arch) // todo extract to push job

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
