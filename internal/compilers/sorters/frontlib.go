package sorters

import (
	"xs/internal/compilers/extractors"
	"xs/internal/compilers/sorters/fl"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_build "xs/internal/jobs/jobs-build"
	jobs_deploy "xs/internal/jobs/jobs-deploy"
	"xs/pkg/io"
)

func NowFrontLibSorter(publish bool) Sorter {
	wm, err := configs.GetInstanceWsManager()
	if err != nil {
		io.Panic(err)
	}
	return &FrontLibSorter{wm: wm, publish: publish}
}

type FrontLibSorter struct {
	wm      *configs.WorkspaceManager
	publish bool
}

func (s *FrontLibSorter) JobCreate(params map[string]string) []jobs.PrintableJob {
	arr := []jobs.PrintableJob{jobs_build.NewBuildFrontLib(params, true)}
	if s.publish {
		arr = append(arr, jobs_deploy.NewDeployFrontLib(params, true))
	}
	return arr
}

func (s *FrontLibSorter) Sort(libs []*configs.XsModule) []jobs.PrintableJob {
	libCompiler := fl.NewLibCompileController(s.wm)
	libCompiler.LoadConfigs(libs)
	return libCompiler.MakeJobs(false, extractors.Frontlib{}, s.JobCreate)
}
