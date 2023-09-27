package compilers

import (
	"xs/internal/compilers/extractors"
	"xs/internal/compilers/sorters"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_build "xs/internal/jobs/jobs-build"
)

type CompilePlanning struct {
	orders map[string]sorters.Sorter
}

func NewCompilePlanning() *CompilePlanning {
	c := &CompilePlanning{}
	c.orders = map[string]sorters.Sorter{}
	FONTLIB := "frontlib"
	BOOTSTRAP := "bootstrap"
	MICROFRONTEND := "microfrontend"
	HELM := "helm"
	CONTAINER := "container"
	//BACKLIB := "backlib"
	c.orders[FONTLIB] = sorters.NowFrontLibSorter()
	c.orders[BOOTSTRAP] = sorters.NewUniversalSorter(jobs_build.NewBuildFrontend, BOOTSTRAP, extractors.Frontend{})
	c.orders[MICROFRONTEND] = sorters.NewUniversalSorter(jobs_build.NewMicroFronted, MICROFRONTEND, extractors.Microfrontend{})
	c.orders[HELM] = sorters.NewUniversalSorter(jobs_build.NewBuildHelm, HELM, extractors.Backend{})
	c.orders[CONTAINER] = sorters.NewUniversalSorter(jobs_build.NewBuildContainer, CONTAINER, extractors.Backend{})
	return c
}

func (c *CompilePlanning) GetPlan(section string, libs []*configs.XsModule) []jobs.PrintableJob {
	return c.orders[section].Sort(libs)
}
