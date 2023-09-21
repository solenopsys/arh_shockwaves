package configs

import (
	"gopkg.in/yaml.v3"
	"reflect"
	"sync"
	"xs/pkg/io"
	"xs/pkg/tools"
)

type ProcessorType string

// Define constants to represent enum values using iota
const (
	PreProcessor  ProcessorType = "pre"
	PostProcessor ProcessorType = "post"
)

type Trigger struct {
	Type     ProcessorType `yaml:"type"`
	Sections []string      `yaml:"sections"`
	Command  []string      `yaml:"command"`
}

type Processor struct {
	Description string     `yaml:"description"`
	Triggers    []*Trigger `yaml:"triggers"`
}

type Git struct {
	Paths    map[string]string `yaml:"paths"`
	Prefixes map[string]string `yaml:"prefixes"`
}

type Jobs struct {
	Builders   map[string][]string  `yaml:"builders"`
	Processors map[string]Processor `yaml:"processors"`
}

type Hosts struct {
	IpfsHost           string `yaml:"ipfsNode"`
	IpfsClusterHost    string `yaml:"ipfsClusterNode"`
	PinningHost        string `yaml:"pinningService"`
	HelmRepositoryHost string `yaml:"helmRepository"`
}

type Configuration struct {
	Hosts     *Hosts                       `yaml:"hosts"`
	Format    string                       `yaml:"format"`
	Git       *Git                         `yaml:"git"`
	Templates map[string]map[string]string `yaml:"templates"`
	Jobs      *Jobs                        `yaml:"jobs"`
}

type ConfigurationManager struct {
	Conf *Configuration
}

func triggerValidate(trigger *Trigger, section string, processorType ProcessorType, command []string) bool {
	triggerTypeOk := trigger.Type == processorType
	commandOk := reflect.DeepEqual(trigger.Command, command)
	var sectionOk bool = false
	for _, currentSection := range trigger.Sections {
		if currentSection == section {
			sectionOk = true
		}
	}
	return triggerTypeOk && commandOk && sectionOk
}

func (m *ConfigurationManager) GetProcessors(section string, processorType ProcessorType, command []string) []string {

	var processorNames = make([]string, 0)

	for name, processor := range m.Conf.Jobs.Processors {

		for _, trigger := range processor.Triggers {
			if triggerValidate(trigger, section, processorType, command) {
				processorNames = append(processorNames, name)
			}
		}

	}

	return processorNames
}

func (m *ConfigurationManager) GetTemplateDirectory(dir string) string {
	return m.Conf.Templates["sections"][dir]
}

func (m *ConfigurationManager) GetBuildersMapping() map[string]string {
	var result = make(map[string]string)
	for builder, sections := range m.Conf.Jobs.Builders {
		for _, section := range sections {
			result[section] = builder
		}
	}
	return result
}

func LoadConfigFile(fileName string) (*Configuration, error) {
	config := &Configuration{}
	fileData, err := tools.ReadFile(fileName)
	if err == nil {
		err = yaml.Unmarshal([]byte(fileData), config)
	} else {
		return nil, err
	}
	return config, err
}

var confInstance *ConfigurationManager
var confOnce sync.Once

func GetInstanceConfManager() *ConfigurationManager {
	confOnce.Do(func() {
		programDir, err := tools.GetProgramDir()
		if err != nil {
			io.Panic(err)
		}
		file, err := LoadConfigFile(programDir + "/xs.config.yaml")
		if err != nil {
			io.Panic(err)
		}
		confInstance = &ConfigurationManager{
			Conf: file,
		}
	})
	return confInstance
}
