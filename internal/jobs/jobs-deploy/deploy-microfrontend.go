package jobs_deploy

import (
	"path/filepath"
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/wrappers"
)

type DeployMicroFrontend struct {
	params map[string]string
}

func (d *DeployMicroFrontend) Execute() *jobs.Result {
	distDir := d.params["dist"]
	name := d.params["name"]
	version := d.params["version"]

	labels := make(map[string]string)
	labels["type"] = "microfrontend"
	labels["name"] = name
	labels["version"] = version
	absoluteDestPath, err := filepath.Abs(distDir)
	cid, err := tools.IpfsPublishDir(absoluteDestPath, labels)
	if err != nil {
		return &jobs.Result{
			Success: false,
			Error:   err,
		}
	}
	p := wrappers.NewPinning()
	ipnsCid, err := p.SmartName(name, cid)

	io.Debug("ipnsCID", ipnsCid)

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
		Key:         "deploy-microfrontend-" + d.params["name"],
	}
}

func NewDeployMicroFrontend(params map[string]string, printConsole bool) jobs.PrintableJob {
	return &DeployMicroFrontend{
		params: params,
	}
}
