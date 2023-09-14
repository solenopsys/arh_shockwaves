package extractors

import (
	"strings"
	"xs/pkg/wrappers"
)

type Frontend struct {
}

func (e Frontend) Extract(name string, path string) map[string]string {
	arr := strings.Split(name, "/")
	dest := wrappers.LoadAngularDest(arr[1], path)
	params := map[string]string{
		"path": path,
		"dest": dest,
		"name": name,
	}
	return params
}
