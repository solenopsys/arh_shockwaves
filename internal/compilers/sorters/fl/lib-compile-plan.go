package fl

import (
	"xs/internal"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type LibCompilePlan struct {
	compileNow    map[string]bool
	packagesOrder *NpmLibPackagesOrder
	xsManager     *configs.WorkspaceManager
	libGroup      string
}

func NewLibCompileController(xm *configs.WorkspaceManager) *LibCompilePlan {
	c := &LibCompilePlan{}
	c.compileNow = map[string]bool{}
	c.xsManager = xm
	return c
}

func (c *LibCompilePlan) LoadConfigs(libs []*configs.XsModule) {

	ord := NewNpmLibPackagesOrder(false)

	for _, lib := range libs {
		lp := configs.LoadNpmLibPackage("./" + lib.Directory + "/package.json")
		ord.AddPackage(lp)
	}

	c.packagesOrder = ord
}

func (c *LibCompilePlan) MakeJobs(force bool, extractor internal.CompileParamsExtractor, jobCreate func(params map[string]string) []jobs.PrintableJob) []jobs.PrintableJob { // todo need refactoring

	cache := NewCompileCache(".xs/compiled")

	var results = []jobs.PrintableJob{}

	for {
		list := c.packagesOrder.NextList()

		if list == nil || len(list) == 0 {
			break
		}

		for _, pack := range list {
			xsPackConf := c.xsManager.ExtractModule(pack.Name)

			path := "." + c.libGroup + "/" + xsPackConf.Directory

			c.compileNow[pack.Name] = true

			var params = map[string]string{}

			params = extractor.Extract(pack.Name, path)

			dest := params["dist"]
			dirExists := xstool.Exists(dest)
			excludeDirs := []string{"node_modules"}
			var hashesOk = false
			if dirExists {
				srcHash, errHash := xstool.HashOfDir(path, excludeDirs)
				if errHash != nil {
					io.Panic(errHash)
				}
				dstHash, errHash := xstool.HashOfDir(dest, excludeDirs)
				if errHash != nil {
					io.Panic(errHash)
				}
				hashesOk = cache.CheckHash(srcHash, dstHash)
			}
			if hashesOk && !force {
				io.PrintColor("SKIP", io.Blue)

				c.packagesOrder.SetCompiled(pack.Name)
				c.compileNow[pack.Name] = false
			} else {

				currentJob := jobCreate(params)

				results = append(results, currentJob...)

				c.compileNow[pack.Name] = false
				c.packagesOrder.SetCompiled(pack.Name)
			}

		}

	}
	return results
}
