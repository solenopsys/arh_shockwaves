package jobs_deploy

import (
	"xs/internal/jobs"
	"xs/pkg/io"
)

type DeployHelm struct {
	params map[string]string
}

func (d *DeployHelm) Execute() *jobs.Result {
	path := d.params["dist"]

	println("not implemented: ", path)

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
