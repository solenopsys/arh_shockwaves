package jobs

import "xs/pkg/io"

type Result struct {
	Success     bool
	Err         error
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
		io.Println(result.Err.Error())
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
