package extractors

import "xs/pkg/wrappers"

type Frontend struct {
}

func (e Frontend) Extract(name string, path string) map[string]string {
	dest := wrappers.LoadAngularDest(name, path)
	params := map[string]string{
		"dest": dest,
		"name": name,
	}
	return params
}
