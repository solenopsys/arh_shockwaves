package extractors

import "xs/pkg/wrappers"

type Frontlib struct {
	Base map[string]string
}

func (e Frontlib) Extract(name string, path string) map[string]string {
	dest := wrappers.LoadNgDest(path)
	params := map[string]string{
		"path": path,
		"dest": dest,
	}

	for k, v := range e.Base {
		params[k] = v
	}
	return params
}
