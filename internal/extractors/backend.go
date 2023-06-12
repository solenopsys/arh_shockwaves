package extractors

type Backend struct {
}

func (e Backend) Extract(name string, path string) map[string]string {
	params := map[string]string{
		"name": name,
	}
	return params
}
