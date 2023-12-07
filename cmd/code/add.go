package code

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
	"xs/internal/services"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/ui"
)

var cmdAdd = &cobra.Command{
	Use:   "add [pattern]",
	Short: "Tags section monorepo",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]

		err := tools.ToWorkspaceRootDir()
		if err != nil {
			io.Fatal("Workspace root dir not found")
		}

		jobsPlan := jobs.ConvertPjToJi(makeAddPlan("*"))
		applied, changedPattern := ui.FilteredListView(jobsPlan, "Sources for loaded", pattern)
		if applied {
			jobsPlanApplied := makeAddPlan(changedPattern)
			executor := jobs.NewExecutor(jobsPlanApplied)
			ui.ProcessingJobs(executor)
		}
	},
}

func makeAddPlan(pattern string) []jobs.PrintableJob {

	templatesJobs := make(map[string]*jobs.PrintableJob)
	codeJobs := make([]jobs.PrintableJob, 0)
	pinning := services.NewPinningRequests()
	repos, err := pinning.FindRepo(pattern)
	if err != nil {
		io.Fatal(err)
	}
	for packageName, val := range *repos {

		directory := val.To
		packPath := strings.Replace(packageName, "@", "/", -1)
		moduleSubDir := directory + packPath
		moduleSubDirExists := tools.Exists(moduleSubDir)
		subDirs := strings.Split(directory, "/")
		var subDir string
		if len(subDirs) > 1 {
			subDir = strings.Join(subDirs[0:2], "/")
		} else {
			subDir = subDirs[0]
		}

		templateJob := checkTemplateExists(subDir)
		templatesJobs[subDir] = &templateJob

		if !moduleSubDirExists {
			var loadJob jobs.PrintableJob
			loadJob = jobs_fetch.NewCodeLoad(val.Cid, packageName, moduleSubDir, val.Src)

			processorsManager := CreateProcessors([]string{"code", "add"})
			preJobs := processorsManager.GetPreProcessors(subDir, packageName, moduleSubDir)
			postJobs := processorsManager.GetPostProcessors(subDir, packageName, moduleSubDir)
			codeJobs = append(codeJobs, preJobs...)
			codeJobs = append(codeJobs, loadJob)
			codeJobs = append(codeJobs, postJobs...)
		}
	}

	for _, val := range templatesJobs {
		if *val != nil {
			codeJobs = append([]jobs.PrintableJob{*val}, codeJobs...)
		}
	}

	return codeJobs
}

func checkTemplateExists(subDir string) jobs.PrintableJob {
	confManager := configs.GetInstanceConfManager()
	subDirExists := tools.Exists(subDir)
	if !subDirExists {
		templateModule := confManager.GetTemplateDirectory(subDir)
		return jobs_fetch.NewTemplateLoad(templateModule, subDir)
	} else {
		return nil
	}
}
