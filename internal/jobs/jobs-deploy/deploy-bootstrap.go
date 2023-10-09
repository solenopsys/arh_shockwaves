package jobs_deploy

import (
	"xs/internal/jobs"
	"xs/pkg/io"
)

type DeployFrontendBootstrap struct {
	params map[string]string
}

func (d *DeployFrontendBootstrap) Execute() *jobs.Result {
	path := d.params["dist"]

	println("not implemented: ", path)

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "BuildHelm executed",
	}
}

func (d *DeployFrontendBootstrap) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Deploy frontend bootstrap " + d.params["dist"],
		Short:       "Reddy",
	}
}

func NewDeployFrontendBootstrap(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployFrontendBootstrap{
		params: params,
	}
}
