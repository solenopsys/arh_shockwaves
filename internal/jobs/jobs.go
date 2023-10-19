package jobs

import (
	"github.com/charmbracelet/lipgloss"
)

type Result struct {
	Success     bool
	Error       error
	Description string
}

type Job interface {
	Execute() *Result
}

type ItemTitle struct {
	Style       lipgloss.Style
	Description string
	Name        string
	Key         string
}

type UiListItem interface {
	Title() ItemTitle
}

type PrintableJob interface {
	Job
	UiListItem
}

func ConvertPjToJi(jobsList []PrintableJob) []ItemTitle {
	var ex []ItemTitle

	for _, printableJob := range jobsList {
		ex = append(ex, printableJob.Title())
	}
	return ex
}

//	func ExecuteOneSync(job Job) {
//		result := (job).Execute()
//		PrintResult(result)
//	}
func ExecuteJobsSync(jobs []Job) []Result {
	var results []Result
	for _, job := range jobs {
		result := job.Execute()
		results = append(results, *result)
	}
	return results
}

//
//func ConvertJobs(jobsList []PrintableJob) []Job {
//	var ex []Job
//
//	for _, printableJob := range jobsList {
//		var job Job = printableJob
//		ex = append(ex, job)
//	}
//	return ex
//}
