package jobs_deploy

import (
	"xs/internal/jobs"
	"xs/pkg/tools"
)

type DeployMicroFrontend struct {
	params map[string]string
}

func (d *DeployMicroFrontend) Execute() *jobs.Result {
	distDir := d.params["dist"]
	name := d.params["name"]

	labels := make(map[string]string)
	labels["type"] = "microfrontend"
	labels["name"] = name
	// labels["version"] = todo

	err := tools.IpfsPublishDir(distDir, labels)

	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	} else {
		return &jobs.Result{
			Success:     true,
			Error:       nil,
			Description: "Microfrontend deployed executed " + name,
		}
	}
}

func (d *DeployMicroFrontend) Title() jobs.ItemTitle {
	return jobs.ItemTitle{
		Style:       jobs.DEFAULT_STYLE,
		Description: "Deploy micro-frontend " + d.params["dist"],
		Name:        d.params["name"],
	}
}

func NewDeployMicroFrontend(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployMicroFrontend{
		params: params,
	}
}
