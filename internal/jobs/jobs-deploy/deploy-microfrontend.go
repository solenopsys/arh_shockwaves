package jobs_deploy

import (
	"xs/internal/jobs"
	"xs/pkg/io"
)

type DeployMicroFrontend struct {
	params map[string]string
}

func (d *DeployMicroFrontend) Execute() *jobs.Result {
	path := d.params["dist"]

	println("not implemented: ", path)

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "BuildHelm executed",
	}
}

func (d *DeployMicroFrontend) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Deploy micro-frontend " + d.params["dist"],
		Short:       "Reddy",
	}
}

func NewDeployMicroFrontend(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployMicroFrontend{
		params: params,
	}
}
