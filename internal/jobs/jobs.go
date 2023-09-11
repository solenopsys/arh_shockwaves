package jobs

import "xs/pkg/io"

type Result struct {
	Success     bool
	Err         error
	Description string
}

type Jobs interface {
	Execute() *Result
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

func ExecuteOneSync(job Jobs) {
	result := job.Execute()
	PrintResult(result)
}

func ExecuteJobsSync(jobs []Jobs) {
	for _, job := range jobs {
		result := job.Execute()
		PrintResult(result)
	}
}
