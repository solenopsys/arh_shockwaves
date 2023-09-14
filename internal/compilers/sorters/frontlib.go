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
	return &FrontLibSorter{}
}

type FrontLibSorter struct {
	wm *configs.WorkspaceManager
}

func (s *FrontLibSorter) JobCreate() jobs.PrintableJob {
	return jobs_build.NewBuildFrontLib(map[string]string{}, false)
}

func (s *FrontLibSorter) Sort(libs []*configs.XsModule) []jobs.PrintableJob {
	io.Println("SetCompiled all libraries")

	//	executor.PrintConsole= false
	libCompiler := fl.NewLibCompileController(s.wm, "packages")
	io.Println("Scan directories")
	libCompiler.LoadConfigs(libs)
	io.Println("Start compile")
	return libCompiler.MakeJobs(false, extractors.Frontlib{}, s.JobCreate)
}
