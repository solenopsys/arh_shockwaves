package code

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
	"xs/pkg/tools"
	"xs/pkg/wrappers"
)

var cmdLoad = &cobra.Command{
	Use:   "load",
	Short: "Tags section monorepo",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]

		jobsPlan := makePlan(pattern)

		for _, job := range jobsPlan {
			fmt.Println((*job).Description())
		}

		confirm := tools.ConfirmDialog("Load packets?")

		if confirm {
			fmt.Println("Proceeding with the action.")

			var ex []jobs.Job

			for _, printableJob := range jobsPlan {
				var job jobs.Job = *printableJob
				ex = append(ex, job)
			}
			jobs.ExecuteJobsSync(ex)

		} else {
			fmt.Println("Canceled.")
		}

	},
}

func makePlan(pattern string) []*jobs.PrintableJob {
	templatesJobs := make(map[string]*jobs.PrintableJob)
	codeJobs := make([]*jobs.PrintableJob, 0)
	pinning := wrappers.NewPinning()
	repos, err := pinning.FindRepo(pattern)
	if err != nil {
		log.Fatal(err)
	}
	for packageName, val := range *repos {
		directory := val.To
		packPath := strings.Replace(packageName, "@", "/", -1)
		moduleSubDir := directory + packPath
		moduleSubDirExists := tools.Exists(moduleSubDir)
		subDir := strings.Split(directory, "/")[0]
		templateJob := checkTemplateExists(subDir)
		templatesJobs[subDir] = &templateJob

		if !moduleSubDirExists {
			var loadJob jobs.PrintableJob
			loadJob = jobs_fetch.NewCodeLoad(val.Cid, packageName, moduleSubDir)
			codeJobs = append(codeJobs, &loadJob)
		} else {
			println("Already loaded ", moduleSubDir)
		}
	}

	for _, val := range templatesJobs {
		if *val != nil {
			codeJobs = append([]*jobs.PrintableJob{val}, codeJobs...)
		}
	}

	return codeJobs
}

func checkTemplateExists(subDir string) jobs.PrintableJob {
	wsManager, err := configs.NewWsManager()
	if err != nil {
		log.Fatal(err)
	}
	subDirExists := tools.Exists(subDir)
	if !subDirExists {
		templateModule := wsManager.GetTemplateDirectory(subDir)
		return jobs_fetch.NewTemplateLoad(templateModule, subDir)
	} else {
		return nil
	}
}
