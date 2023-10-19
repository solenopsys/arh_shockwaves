package jobs_deploy

import (
	"xs/internal/jobs"
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

func (d *DeployFrontendBootstrap) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Deploy frontend bootstrap " + d.params["dist"],
		Name:        d.params["name"],
		Key:         "deploy-bootstrap-" + d.params["name"],
	}
}

func NewDeployFrontendBootstrap(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployFrontendBootstrap{
		params: params,
	}
}
