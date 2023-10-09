package jobs_deploy

import (
	"xs/internal/jobs"
	"xs/pkg/io"
)

type DeployContainer struct {
	params map[string]string
}

func (d *DeployContainer) Execute() *jobs.Result {
	path := d.params["dist"]

	println("not implemented: ", path)

	return &jobs.Result{
		Success:     true,
		Error:       nil,
		Description: "BuildHelm executed",
	}
}

func (d *DeployContainer) Description() jobs.JobDescription {
	return jobs.JobDescription{
		Color:       io.Blue,
		Description: "Deploy container " + d.params["dist"],
		Short:       "Reddy",
	}
}

func NewDeployContainer(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployContainer{
		params: params,
	}
}
