package jobs_deploy

import (
	"path/filepath"
	"xs/internal/jobs"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/wrappers"
)

type DeployFrontendBootstrap struct {
	params map[string]string
}

func (d *DeployFrontendBootstrap) Execute() *jobs.Result {
	distDir := d.params["dist"]
	name := d.params["name"]

	labels := make(map[string]string)
	labels["type"] = "bootstrap"
	labels["name"] = name
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
			Description: "Bootstap deployed executed " + name,
		}
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
