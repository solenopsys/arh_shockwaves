package sorters

import (
	"xs/internal"
	"xs/internal/configs"
	"xs/internal/jobs"
	"xs/pkg/io"
)

func NewUniversalSorter(jobCreate func(params map[string]string, printConsole bool) jobs.PrintableJob) Sorter {
	wm, err := configs.GetInstanceWsManager()
	if err != nil {
		io.Panic(err)
	}
	return &UniversalSorter{wm: wm, jobCreate: jobCreate}
}

type UniversalSorter struct {
	wm        *configs.WorkspaceManager
	extractor internal.CompileExecutor
	jobCreate func(params map[string]string, printConsole bool) jobs.PrintableJob
}

func (s *UniversalSorter) JobCreate(params map[string]string) jobs.PrintableJob {
	return s.jobCreate(params, true)
}

func SortByName(libs []*configs.XsModule) { // todo move to tools
	for i := 0; i < len(libs); i++ {
		for j := i + 1; j < len(libs); j++ {
			if libs[i].Name > libs[j].Name {
				libs[i], libs[j] = libs[j], libs[i]
			}
		}
	}
}

func (s *UniversalSorter) Sort(libs []*configs.XsModule) []jobs.PrintableJob {
	SortByName(libs)
	result := []jobs.PrintableJob{}
	for _, lib := range libs {
		result = append(result, s.JobCreate(map[string]string{"name": lib.Name, "dir": lib.Directory}))
	}
	return result
}
