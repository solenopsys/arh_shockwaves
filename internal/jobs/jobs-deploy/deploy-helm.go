package jobs_deploy

import (
	"os"
	"path/filepath"
	"xs/internal/jobs"
	"xs/pkg/wrappers"
)

type DeployHelm struct {
	params map[string]string
}

func (d *DeployHelm) Execute() *jobs.Result {
	dist := d.params["dist"]
	path := d.params["path"]

	parent := filepath.Base(path)
	distFile := dist + "/" + parent + ".tar.gz"
	archBytes, err := os.ReadFile(distFile)

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}

	resp, err := wrappers.PushDir(archBytes)

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	} else {
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "Build Helm executed" + resp,
		}
	}
}

func (d *DeployHelm) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Deploy Helm " + d.params["path"],
		Name:        d.params["name"],
		Key:         "deploy-name-" + d.params["name"],
	}
}

func NewDeployHelm(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployHelm{
		params: params,
	}
}
