package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"xs/pkg/io"
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

type ConfigurationManager struct {
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

	processors := map[string]Processor{}

	err := viper.UnmarshalKey("jobs.processors", &processors)
	if err != nil {
		io.Panic(err)
	}
	for name, processor := range processors {

		for _, trigger := range processor.Triggers {
			if triggerValidate(trigger, section, processorType, command) {
				processorNames = append(processorNames, name)
			}
		}

	}

	return processorNames
}

func (m *ConfigurationManager) GetTemplateDirectory(dir string) string {
	return viper.GetString("templates.sections." + dir)
}

func (m *ConfigurationManager) GetBuildersMapping() map[string]string {
	var result = make(map[string]string)
	builders := map[string][]string{}

	err := viper.UnmarshalKey("jobs.builders", &builders)
	if err != nil {
		io.Panic(err)
	}
	for builder, sections := range builders {
		for _, section := range sections {
			result[section] = builder
		}
	}
	return result
}

var confInstance *ConfigurationManager
var confOnce sync.Once

const XS_CONFIGURATION_FILE = "xs-configuration"

func GetInstanceConfManager() *ConfigurationManager {
	confOnce.Do(func() {
		programDir, err := GetProgramDir()
		if err != nil {
			io.Panic(err)
		}
		viper.SetConfigName(XS_CONFIGURATION_FILE) // name of config file (without extension)
		viper.SetConfigType("yaml")                // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(programDir)            // path to look for the config file in
		viper.AddConfigPath(".")                   // optionally look for config in the working directory
		err = viper.ReadInConfig()                 // Find and read the config file
		if err != nil {                            // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		if err != nil {
			io.Panic(err)
		}
		confInstance = &ConfigurationManager{}
	})
	return confInstance
}

func GetProgramDir() (string, error) {
	// Get the path to the executable binary
	exePath, err := os.Executable()

	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)

	return exeDir, nil
}
