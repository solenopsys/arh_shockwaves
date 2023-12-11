package extractors

import "strings"

type Bootstrap struct {
}

func (e Bootstrap) Extract(name string, path string) map[string]string {

	packageName := strings.Replace(name, "@", "", 1)
	distribution := "./converged/dist/bootstraps/" + packageName

	params := map[string]string{
		"path": path,
		"dist": distribution,
		"name": name,
	}
	return params
}
