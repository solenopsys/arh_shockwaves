package services

import "xs/utils"

type NpmLibPackagesOrder struct {
	packages map[string]*utils.NpmLibPackage
	compiled map[string]bool
	verbose  bool
}

func (o *NpmLibPackagesOrder) filterDeps(pack *utils.NpmLibPackage) []string {

	var res = []string{}

	for dep, _ := range pack.Dependencies {
		if o.packages[dep] != nil {
			res = append(res, dep)
		}
	}
	for dep, _ := range pack.PeerDependencies {
		if o.packages[dep] != nil {
			res = append(res, dep)
		}
	}
	for dep, _ := range pack.AllowedNonPeerDependencies {
		if o.packages[dep] != nil {
			res = append(res, dep)
		}
	}

	return res
}

func (o *NpmLibPackagesOrder) count() int {
	return len(o.packages)
}

func (o *NpmLibPackagesOrder) NextList() []*utils.NpmLibPackage {
	var m = map[string]*utils.NpmLibPackage{}
	for _, p := range o.packages {
		if !o.compiled[p.Name] {
			if o.verbose {
				println("pack name", p.Name)
			}

			filtered := o.filterDeps(p)

			needCompiledCount := len(filtered)
			if o.verbose {
				println("need compile", needCompiledCount)
			}

			var actualCompiledCount = 0
			if needCompiledCount > 0 {
				for _, dep := range filtered {
					if o.verbose {
						println("\t", dep, " - ", o.compiled[dep])
					}
					if o.compiled[dep] {

						actualCompiledCount++
					}
				}
			}

			if actualCompiledCount == needCompiledCount {
				m[p.Name] = p
			}
		}
	}

	var result = []*utils.NpmLibPackage{}
	for _, p := range m {
		result = append(result, p)
	}

	return result
}

func (o *NpmLibPackagesOrder) SetCompiled(name string) {
	o.compiled[name] = true
}

func (o *NpmLibPackagesOrder) CompileList(list []*utils.NpmLibPackage) {
	for _, p := range list {
		o.SetCompiled(p.Name)
	}
}

func (o *NpmLibPackagesOrder) AddPackage(p *utils.NpmLibPackage) {
	o.packages[p.Name] = p
}

func NewNpmLibPackagesOrder(verbose bool) *NpmLibPackagesOrder {
	return &NpmLibPackagesOrder{
		packages: map[string]*utils.NpmLibPackage{},
		compiled: map[string]bool{},
		verbose:  verbose,
	}
}
