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
	FRONT := "front"
	MODULE := "module"
	HELM := "helm"
	CONTAINER := "container"
	c.orders[FONTLIB] = sorters.NowFrontLibSorter()
	c.orders[FRONT] = sorters.NewUniversalSorter(jobs_build.NewBuildFrontend, FRONT, extractors.Frontend{})
	c.orders[MODULE] = sorters.NewUniversalSorter(jobs_build.NewMicroFronted, MODULE, extractors.Microfrontend{})
	c.orders[HELM] = sorters.NewUniversalSorter(jobs_build.NewBuildHelm, HELM, extractors.Backend{})
	c.orders[CONTAINER] = sorters.NewUniversalSorter(jobs_build.NewBuildContainer, CONTAINER, extractors.Backend{})
	return c
}

func (c *CompilePlanning) GetPlan(section string, libs []*configs.XsModule) []jobs.PrintableJob {
	return c.orders[section].Sort(libs)
}
