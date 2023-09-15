package compilers

import (
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
	c.orders["frontlib"] = sorters.NowFrontLibSorter()
	c.orders["front"] = sorters.NewUniversalSorter(jobs_build.NewBuildFrontend)
	c.orders["module"] = sorters.NewUniversalSorter(jobs_build.NewMicroFronted)
	c.orders["helm"] = sorters.NewUniversalSorter(jobs_build.NewBuildHelm)
	c.orders["container"] = sorters.NewUniversalSorter(jobs_build.NewBuildContainer)
	return c
}

func (c *CompilePlanning) GetPlan(section string, libs []*configs.XsModule) []jobs.PrintableJob {
	return c.orders[section].Sort(libs)
}
