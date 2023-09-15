package sorters

import (
	"xs/internal/compilers/extractors"
	"xs/internal/compilers/sorters/fl"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_build "xs/internal/jobs/jobs-build"
	"xs/pkg/io"
)

func NowFrontLibSorter() Sorter {
	wm, err := configs.GetInstanceWsManager()
	if err != nil {
		io.Panic(err)
	}
	return &FrontLibSorter{wm: wm}
}

type FrontLibSorter struct {
	wm *configs.WorkspaceManager
}

func (s *FrontLibSorter) JobCreate(params map[string]string) jobs.PrintableJob {
	return jobs_build.NewBuildFrontLib(params, true)
}

func (s *FrontLibSorter) Sort(libs []*configs.XsModule) []jobs.PrintableJob {
	libCompiler := fl.NewLibCompileController(s.wm)
	libCompiler.LoadConfigs(libs)
	return libCompiler.MakeJobs(false, extractors.Frontlib{}, s.JobCreate)
}
