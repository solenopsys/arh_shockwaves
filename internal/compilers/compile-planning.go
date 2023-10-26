package compilers

import (
	"xs/internal"
	"xs/internal/compilers/extractors"
	"xs/internal/compilers/sorters"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_build "xs/internal/jobs/jobs-build"
	jobs_deploy "xs/internal/jobs/jobs-deploy"
)

type CompilePlanning struct {
	orders  map[string]sorters.Sorter
	publish bool
}

func (c *CompilePlanning) createUniversal(key string,
	build func(params map[string]string, printConsole bool) jobs.PrintableJob,
	deploy func(params map[string]string, printConsole bool) jobs.PrintableJob,
	extractor internal.CompileParamsExtractor) {

	bootstrapJobs := []func(params map[string]string, printConsole bool) jobs.PrintableJob{
		build,
	}
	if c.publish {
		bootstrapJobs = append(bootstrapJobs, deploy)
	}
	c.orders[key] = sorters.NewUniversalSorter(
		bootstrapJobs,
		key,
		extractor,
	)
}

func NewCompilePlanning(publish bool) *CompilePlanning {
	c := &CompilePlanning{publish: publish}
	c.orders = map[string]sorters.Sorter{}
	c.orders[("frontlib")] = sorters.NowFrontLibSorter(publish)
	c.createUniversal("bootstrap", jobs_build.NewBuildFrontend, jobs_deploy.NewDeployFrontendBootstrap, extractors.Bootstrap{})
	c.createUniversal("microfrontend", jobs_build.NewMicroFronted, jobs_deploy.NewDeployMicroFrontend, extractors.Microfrontend{})
	c.createUniversal("helm", jobs_build.NewBuildHelm, jobs_deploy.NewDeployHelm, extractors.Backend{})
	c.createUniversal("container", jobs_build.NewBuildContainer, jobs_deploy.NewDeployContainer, extractors.Backend{})
	return c
}

func (c *CompilePlanning) GetPlan(section string, libs []*configs.XsModule) []jobs.PrintableJob {
	return c.orders[section].Sort(libs)
}
