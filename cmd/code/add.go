package code

import (
	"github.com/spf13/cobra"
	"strings"
	"xs/internal/configs"
	"xs/internal/jobs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
	"xs/pkg/io"
	"xs/pkg/tools"
	"xs/pkg/wrappers"
)

var cmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Tags section monorepo",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]

		err := tools.ToWorkspaceRootDir()
		if err != nil {
			io.Fatal("Workspace root dir not found")
		}

		jobsPlan := makePlan(pattern)

		for _, job := range jobsPlan {
			io.Println(job.Description())
		}

		confirm := tools.ConfirmDialog("Load packets?")

		if confirm {
			io.Println("Proceeding with the action.")
			jobs.ExecuteJobsSync(jobs.ConvertJobs(jobsPlan))

		} else {
			io.Println("Canceled.")
		}

	},
}

func makePlan(pattern string) []jobs.PrintableJob {
	confManager := configs.GetInstanceConfManager()

	templatesJobs := make(map[string]*jobs.PrintableJob)
	codeJobs := make([]jobs.PrintableJob, 0)
	pinning := wrappers.NewPinning()
	repos, err := pinning.FindRepo(pattern)
	if err != nil {
		io.Fatal(err)
	}
	for packageName, val := range *repos {

		directory := val.To
		packPath := strings.Replace(packageName, "@", "/", -1)
		moduleSubDir := directory + packPath
		moduleSubDirExists := tools.Exists(moduleSubDir)
		subDir := strings.Split(directory, "/")[0]
		templateJob := checkTemplateExists(subDir)

		processorsMapping := make(map[string]jobs.PrintableJob) // todo export to global constant
		processorsMapping["ts_injector"] = jobs_fetch.NewTsConfigModuleInject(packageName, moduleSubDir)

		templatesJobs[subDir] = &templateJob

		if !moduleSubDirExists {
			var loadJob jobs.PrintableJob
			loadJob = jobs_fetch.NewCodeLoad(val.Cid, packageName, moduleSubDir, val.Src)
			preJobs := processingJobs(*confManager, configs.PreProcessor, subDir, processorsMapping)
			postJobs := processingJobs(*confManager, configs.PostProcessor, subDir, processorsMapping)
			codeJobs = append(codeJobs, preJobs...)
			codeJobs = append(codeJobs, loadJob)
			codeJobs = append(codeJobs, postJobs...)
		} else {
			io.Println("Already loaded ", moduleSubDir)
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

func processingJobs(
	confManager configs.ConfigurationManager,
	processorType configs.ProcessorType,
	subDir string,
	processorsMapping map[string]jobs.PrintableJob) []jobs.PrintableJob {

	processorsJobs := make([]jobs.PrintableJob, 0)

	processorsNames := confManager.GetProcessors(subDir, processorType, []string{"code", "add"})

	for _, processorName := range processorsNames {
		job := processorsMapping[processorName]
		processorsJobs = append(processorsJobs, job)
	}
	return processorsJobs
}
