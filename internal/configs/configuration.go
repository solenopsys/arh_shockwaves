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
	Type     ProcessorType `json:"type"`
	Sections []string      `json:"sections"`
	Command  []string      `json:"command"`
}

type Processor struct {
	Description string     `json:"description"`
	Triggers    []*Trigger `json:"triggers"`
}

type Configuration struct {
	Format     string                       `json:"format"`
	Templates  map[string]map[string]string `json:"templates"`
	Builders   map[string][]string          `json:"builders"`
	Processors map[string]Processor         `json:"processors"`
}

type ConfigurationManager struct {
	configuration *Configuration
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

	for name, processor := range m.configuration.Processors {

		for _, trigger := range processor.Triggers {
			if triggerValidate(trigger, section, processorType, command) {
				processorNames = append(processorNames, name)
			}
		}

	}

	return processorNames
}

func (m *ConfigurationManager) GetTemplateDirectory(dir string) string {
	return m.configuration.Templates["sections"][dir]
}

func (m *ConfigurationManager) GetBuildersMapping() map[string]string {
	var result = make(map[string]string)
	for builder, sections := range m.configuration.Builders {
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

func GetInstanceConfManager() (*ConfigurationManager, error) {
	confOnce.Do(func() {
		file, err := LoadConfigFile("./configuration.yaml")
		if err != nil {
			io.Panic(err)
		}
		confInstance = &ConfigurationManager{
			configuration: file,
		}
	})
	return confInstance, nil
}
