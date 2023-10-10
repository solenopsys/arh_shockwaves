package jobs_build

import (
	"os"
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
	dist := b.params["dist"]
	pathHelmDir := path + "/helm"

	parent := filepath.Base(path)
	os.Mkdir(dist, 0777)

	distFile := dist + "/" + parent + ".tar.gz"
	absolutPath, _ := filepath.Abs(pathHelmDir)
	io.Println("path", absolutPath)

	arch := wrappers.ArchiveDir(absolutPath, "helm")

	io.Println("archive size", len(arch))

	err := os.WriteFile(distFile, arch, 0444)

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "BuildHelm executed: " + distFile,
	}
}

func (b *BuildHelm) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Build helm " + b.params["path"],
		Short:       "Reddy",
	}
}

func NewBuildHelm(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &BuildHelm{
		params: params,
	}
}
