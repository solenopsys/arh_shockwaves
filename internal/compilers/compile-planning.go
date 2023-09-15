package compilers

import (
	"xs/internal/compilers/sorters"
	"xs/internal/configs"
	"xs/internal/jobs"
)

type CompilePlanning struct {
	orders map[string]sorters.Sorter
}

func NewCompilePlanning() *CompilePlanning {
	c := &CompilePlanning{}
	c.orders = map[string]sorters.Sorter{}
	c.orders["frontlib"] = sorters.NowFrontLibSorter()
	return c

}

func (c *CompilePlanning) GetPlan(section string, libs []*configs.XsModule) []jobs.PrintableJob {
	return c.orders[section].Sort(libs)
}
