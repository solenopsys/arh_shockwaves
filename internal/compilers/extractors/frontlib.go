package extractors

import "xs/pkg/wrappers"

type Frontlib struct {
	Base map[string]string
}

func (e Frontlib) Extract(name string, path string) map[string]string {
	distribution := wrappers.LoadNgDest(path)
	params := map[string]string{
		"name": name,
		"path": path,
		"dist": distribution,
	}

	for k, v := range e.Base {
		params[k] = v
	}
	return params
}
