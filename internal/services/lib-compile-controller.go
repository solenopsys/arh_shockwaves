package services

import (
	"strconv"
	"xs/internal"
	"xs/internal/configs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
)

type LibCompileController struct {
	compileNow      map[string]bool
	packagesOrder   *NpmLibPackagesOrder
	compileExecutor internal.CompileExecutor
	xsManager       *configs.XsManager
}

func NewLibCompileController(xm *configs.XsManager, executor internal.CompileExecutor) *LibCompileController {
	c := &LibCompileController{}
	c.compileNow = map[string]bool{}
	c.compileExecutor = executor
	c.xsManager = xm
	return c
}

func (c *LibCompileController) LoadPlan(libGroup string, libs []*configs.XsMonorepoModule) {

	ord := NewNpmLibPackagesOrder(false)

	for _, lib := range libs {
		lp := configs.LoadNpmLibPackage("./" + libGroup + "/" + lib.Directory + "/package.json")
		ord.AddPackage(lp)
	}

	c.packagesOrder = ord
}

func (c *LibCompileController) CompileOnOneThread(force bool, libGroup string, extractor internal.CompileParamsExtractor) { // todo need refactoring

	cache := NewCompileCache(".xs/compiled")

	io.Println("Packages count: ", c.packagesOrder.count())
	var n = 0
	for {
		list := c.packagesOrder.NextList()

		if list == nil || len(list) == 0 {
			break
		}

		for _, pack := range list {
			xsPackConf := c.xsManager.ExtractModule(libGroup, pack.Name)

			path := libGroup + "/" + xsPackConf.Directory
			n++
			strN := strconv.Itoa(n)
			io.Print(strN + " : " + xsPackConf.Npm)
			c.compileNow[pack.Name] = true

			var params = map[string]string{}

			if extractor != nil {
				params = extractor.Extract(pack.Name, path)
			}

			dest := params["dest"]
			dirExists := xstool.DirExists(dest)
			excludeDirs := []string{"node_modules"}
			var hashesOk = false
			if dirExists { //todo refactor

				srcHash, errHash := xstool.HashOfDir(path, excludeDirs)
				if errHash != nil {
					io.Panic(errHash)
				}
				dstHash, errHash := xstool.HashOfDir(dest, excludeDirs)
				if errHash != nil {
					io.Panic(errHash)
				}
				hashesOk = cache.checkHash(srcHash, dstHash)
			}
			if hashesOk && !force {
				io.PrintColor("SKIP", io.Blue)

				c.packagesOrder.SetCompiled(pack.Name)
				c.compileNow[pack.Name] = false
			} else {

				err := c.compileExecutor.Compile(params)
				if err != nil {
					io.Panic(err)
				} else {
					srcHash, errHash := xstool.HashOfDir(path, excludeDirs)
					if errHash != nil {
						io.Panic(errHash)
					}
					dstHash, errHash := xstool.HashOfDir(dest, excludeDirs)
					if errHash != nil {
						io.Panic(errHash)
					}
					errHash = cache.saveHash(srcHash, dstHash)
					if errHash != nil {
						io.Panic(errHash)
					}
					c.compileNow[pack.Name] = false
					c.packagesOrder.SetCompiled(pack.Name)
				}
			}

		}

	}
}
