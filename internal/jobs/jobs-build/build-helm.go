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
		Err:         nil,
		Description: "BuildHelm executed",
	}
}
