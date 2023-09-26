package jobs

import (
	"xs/internal/configs"
	jobs_fetch "xs/internal/jobs/jobs-fetch"
)

const TS_INJECTOR = "ts_injector"
const TS_REMOVER = "ts_remover"

type Processors struct {
	mappingCreators map[string]func(packageName string, targetDir string) PrintableJob
	confManager     configs.ConfigurationManager
	commands        []string
}

func (p *Processors) init() {
	p.mappingCreators[TS_INJECTOR] = jobs_fetch.NewTsConfigModuleInject
	p.mappingCreators[TS_REMOVER] = jobs_fetch.NewTsConfigModuleRemove
}

func (p *Processors) GetPreProcessors(subDir string, packageName string, targetDir string) []PrintableJob {
	return p.processingJobs(configs.PreProcessor, subDir, packageName, targetDir)
}

func (p *Processors) GetPostProcessors(subDir string, packageName string, targetDir string) []PrintableJob {
	return p.processingJobs(configs.PreProcessor, subDir, packageName, targetDir)
}

func (p *Processors) processingJobs(processorType configs.ProcessorType, subDir string, packageName string, targetDir string) []PrintableJob {
	processorsJobs := make([]PrintableJob, 0)
	processorsNames := p.confManager.GetProcessors(subDir, processorType, p.commands) //[]string{"code", "add"}

	for _, processorName := range processorsNames {
		creator := p.mappingCreators[processorName]
		job := creator(packageName, targetDir)
		processorsJobs = append(processorsJobs, job)
	}

	return processorsJobs
}

func NewProcessors(commands []string) *Processors {
	confManager := configs.GetInstanceConfManager()
	processors := &Processors{
		mappingCreators: make(map[string]func(packageName string, targetDir string) PrintableJob),
		confManager:     *confManager,
		commands:        commands,
	}
	processors.init()
	return processors
}
