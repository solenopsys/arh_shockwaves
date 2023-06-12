package extractors

import "xs/pkg/wrappers"

type Frontlib struct {
}

func (e Frontlib) Extract(name string, path string) map[string]string {
	dest := wrappers.LoadNgDest(path)
	params := map[string]string{
		"path": path,
		"dest": dest,
	}
	return params
}
