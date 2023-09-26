package code

import (
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
)

func CreateProcessors(command []string) *jobs.Processors { // todo move it
	mappingCreators := map[string]func(packageName string, targetDir string) jobs.PrintableJob{
		"ts_injector": jobs_fetch.NewTsConfigModuleInject,
		"ts_remover":  jobs_fetch.NewTsConfigModuleRemove,
	}
	return jobs.NewProcessors(mappingCreators, command)

}
