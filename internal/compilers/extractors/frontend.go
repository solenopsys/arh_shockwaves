package extractors

import (
	"strings"
	"xs/pkg/wrappers"
)

type Frontend struct {
}

func (e Frontend) Extract(name string, path string) map[string]string {
	arr := strings.Split(name, "/")
	distribution := wrappers.LoadAngularDest(arr[1], path)
	params := map[string]string{
		"path": path,
		"dist": distribution,
		"name": name,
	}
	return params
}
