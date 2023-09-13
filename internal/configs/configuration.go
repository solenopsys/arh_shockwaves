package configs

import (
	"encoding/json"
	"reflect"
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
	Format     Format
	Builders   map[string][]string  `json:"builders"`
	Processors map[string]Processor `json:"processors"`
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

func LoadConfigFile(fileName string) *Configuration {
	config := &Configuration{}
	fileData, err := tools.ReadFile(fileName)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), config)
	} else {
		io.Fatal(err)
	}
	return config
}

func NewConfigurationManager() *ConfigurationManager {
	return &ConfigurationManager{
		configuration: LoadConfigFile("./xs-configuration.json"),
	}
}