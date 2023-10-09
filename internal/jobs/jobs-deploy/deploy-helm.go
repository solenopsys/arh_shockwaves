package jobs_deploy

import (
	"os"
	"path/filepath"
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/wrappers"
)

type DeployHelm struct {
	params map[string]string
}

func (d *DeployHelm) Execute() *jobs.Result {
	dist := d.params["dist"]
	path := d.params["path"]
	parent := filepath.Dir(path)
	distFile := dist + "/" + parent + ".tar.gz"
	archBytes, err := os.ReadFile(distFile)

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	wrappers.PushDir(archBytes)

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "BuildHelm executed",
	}
}

func (d *DeployHelm) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Deploy Helm " + d.params["dist"],
		Short:       "Reddy",
	}
}

func NewDeployHelm(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployHelm{
		params: params,
	}
}
