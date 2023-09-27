package jobs

import (
	"xs/pkg/io"
)

type Result struct {
	Success     bool
	Error       error
	Description string
}

type Job interface {
	Execute() *Result
}

type JobDescription struct {
	Color       io.PrintStyle
	Description string
	Short       string
}

func PrintJob(message JobDescription) {
	io.PrintColor(message.Short, message.Color)
	io.Println(message.Description)
}

type JobInfo interface {
	Description() JobDescription
}

type PrintableJob interface {
	Job
	JobInfo
}

func PrintResult(result *Result) {
	if result == nil {
		io.PrintColor("ERROR", io.Red)
		io.Println("Result is nil")
		return
	}
	if result.Success {
		io.PrintColor("OK", io.Green)
		io.Println(result.Description)
	} else {
		io.PrintColor("ERROR", io.Red)
		io.Println(result.Error.Error())
	}
}

func ExecuteOneSync(job Job) {
	result := (job).Execute()
	PrintResult(result)
}

func ExecuteJobsSync(jobs []Job) {
	for _, job := range jobs {
		result := job.Execute()
		PrintResult(result)
	}
}

func ConvertJobs(jobsList []PrintableJob) []Job {
	var ex []Job

	for _, printableJob := range jobsList {
		var job Job = printableJob
		ex = append(ex, job)
	}
	return ex
}
