package services

import "xs/utils"

type NpmLibPackagesOrder struct {
	packages map[string]*utils.NpmLibPackage
	compiled map[string]bool
}

func (o *NpmLibPackagesOrder) filterDeps(deps map[string]string) []string {

	var res = []string{}

	for dep, _ := range deps {
		if o.packages[dep] != nil {
			res = append(res, dep)
		}
	}

	return res
}

func (o *NpmLibPackagesOrder) NextList() []*utils.NpmLibPackage {
	var m = map[string]*utils.NpmLibPackage{}
	for _, p := range o.packages {
		if !o.compiled[p.Name] {
			filtred := o.filterDeps(p.AllowedNonPeerDependencies)
			needCompiledCount := len(filtred)
			var actualCompiledCount = 0
			if needCompiledCount > 0 {
				for _, dep := range filtred {
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

func (o *NpmLibPackagesOrder) Compile(name string) {
	o.compiled[name] = true
}

func (o *NpmLibPackagesOrder) CompileList(list []*utils.NpmLibPackage) {
	for _, p := range list {
		o.Compile(p.Name)
	}
}

func (o *NpmLibPackagesOrder) AddPackage(p *utils.NpmLibPackage) {
	o.packages[p.Name] = p
}

func NewNpmLibPackagesOrder() *NpmLibPackagesOrder {
	return &NpmLibPackagesOrder{
		packages: map[string]*utils.NpmLibPackage{},
		compiled: map[string]bool{},
	}
}
