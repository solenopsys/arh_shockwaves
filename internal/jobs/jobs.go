package jobs

import "xs/pkg/io"

type Result struct {
	Success     bool
	Error       error
	Description string
}

type Job interface {
	Execute() *Result
}

type JobInfo interface {
	Description() string
}

type PrintableJob interface {
	Job
	JobInfo
}

func PrintResult(result *Result) {
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
