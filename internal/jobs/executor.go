package jobs

type Executor struct {
	results map[string]bool
	index   int
	jobs    []PrintableJob
}

func (j *Executor) JobsCount() int {
	return len(j.jobs)
}

func (j *Executor) SuccessCount() int {
	return j.calcCount(true)
}

func (j *Executor) calcCount(result bool) int {
	successCount := 0
	for _, v := range j.results {
		if v == result {
			successCount++
		}
	}
	return successCount
}

func (j *Executor) ErrorCount() int {
	return j.calcCount(false)
}

func (j *Executor) GetCurrent() PrintableJob {
	return j.jobs[j.index]
}
func (j *Executor) GetCurrentIndex() int {
	return j.index
}

func (j *Executor) Progress() float64 {
	return float64(j.index) / float64(len(j.jobs)-1)
}

func (j *Executor) Next() {
	j.index++
}

func (j *Executor) IsDone() bool {
	return j.index >= len(j.jobs)-1
}

func (j *Executor) RunJob() *Result {
	job := j.GetCurrent()
	execute := job.Execute()
	if execute == nil {
		j.results[job.Title().Key] = false
	} else {
		j.results[job.Title().Key] = execute.Success
	}

	return execute
}

func (j *Executor) GetFailKeys() []string {
	keys := make([]string, 0)
	for _, job := range j.jobs {
		if !j.results[job.Title().Key] {
			keys = append(keys, job.Title().Key)
		}
	}
	return keys
}

func NewExecutor(jobs []PrintableJob) *Executor {
	return &Executor{jobs: jobs, index: 0, results: make(map[string]bool)}
}
