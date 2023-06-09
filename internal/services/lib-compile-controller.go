package services

import (
	"xs/internal/compilers"
	"xs/internal/configs"
	"xs/pkg/io"
	xstool "xs/pkg/tools"
	"xs/pkg/wrappers"
)

func NewLibCompileController(xsFile string, libGroup string) *LibCompileController {
	c := &LibCompileController{}
	c.xsFile = xsFile
	c.libGroup = libGroup
	c.compileNow = map[string]bool{}
	c.compileExecutor = compilers.AngularPackageCompileExecutor{PrintConsole: false}
	c.LoadPlan()
	return c
}

type LibCompileController struct {
	xsFile          string
	libGroup        string
	compileNow      map[string]bool
	packagesOrder   *NpmLibPackagesOrder
	compileExecutor compilers.CompileExecutor
	xsManager       *configs.XsManager
}

func (c *LibCompileController) LoadPlan() {
	fileName := c.xsFile
	xm := &configs.XsManager{}
	c.xsManager = xm
	err := xm.Load(fileName)
	if err != nil {
		io.Panic(err)
	}

	libs := xm.ExtractGroup(c.libGroup)
	ord := NewNpmLibPackagesOrder(false)

	for _, lib := range libs {
		lp := configs.LoadNpmLibPackage("./" + c.libGroup + "/" + lib.Directory + "/package.json")
		ord.AddPackage(lp)
	}

	c.packagesOrder = ord
}

func (c *LibCompileController) CompileOnOneThread(force bool) { //todo need refactoring

	cache := NewCompileCache(".xs/compiled")

	io.Println("Packages count: ", c.packagesOrder.count())
	var n = 0
	for {
		list := c.packagesOrder.NextList()

		if list == nil || len(list) == 0 {
			break
		}
		//io.Println("COMPILE GROUP: ")
		//for _, pack := range list {
		//	io.Println("\t", pack.Name+" ")
		//}

		for _, pack := range list {
			xsPackConf := c.xsManager.Extract(c.libGroup, pack.Name)

			if xsPackConf == nil {
				io.Panic(pack.Name + " not found in " + c.xsFile + " group " + c.libGroup)
			}

			path := c.libGroup + "/" + xsPackConf.Directory
			n++
			io.PrintColor(string(rune(n))+" : "+xsPackConf.Npm+" ", io.Blue)
			c.compileNow[pack.Name] = true

			dest := wrappers.LoadNgDest(path)

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
				params := map[string]string{
					"path": path,
					"dest": dest,
				}
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
