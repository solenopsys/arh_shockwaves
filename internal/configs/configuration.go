package configs

import (
	"encoding/json"
	"xs/pkg/io"
	"xs/pkg/tools"
)

type Processor struct {
	Type      string   `json:"type"`
	Sections  []string `json:"sections"`
	Command   []string `json:"command"`
	Processor string   `json:"processor"`
}

type Configuration struct {
	Format     Format
	Builders   map[string]map[string][]string `json:"builders"`
	Processors []*Processor                   `json:"processors"`
}

type ConfigurationManager struct {
	configuration *Configuration
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
