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

// todo need recovery
//mod, extractError := configs.ExtractModule(name, u.GroupDir, u.RepoType)
//if extractError != nil {
//	io.Panic(extractError)
//}
//
//compiler := u.Executor
//if mod == nil {
//	io.Panic("Module not found")
//}
//io.Println("Mod ", mod.Directory)
//
//path := "./" + u.GroupDir + "/" + mod.Directory
//io.Println("SetCompiled library", path)
//
//var params = map[string]string{}
//
//if u.Extractor != nil {
//	params = u.Extractor.Extract(name, path)
//}
//
//err := compiler.Compile(params)
//
//if err != nil {
//	io.Panic(err)
//}
