package configs

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"sync"
	"xs/pkg/io"
)

type XsModule struct {
	Directory string
	Name      string
}

type Workspace struct {
	Format string                       `yaml:"format"`
	Code   map[string]map[string]string `yaml:"code"`
}

type WorkspaceManager struct {
	workspace *Workspace
	file      string
}

func (m *WorkspaceManager) Load() error {
	fileData, err := os.ReadFile(m.file)
	if err == nil {
		err = yaml.Unmarshal([]byte(fileData), m.workspace)
		return nil
	} else {
		return err
	}
}

func (m *WorkspaceManager) Save() {
	bytes, err := yaml.Marshal(m.workspace)
	if err != nil {
		io.Panic(err)
	} else {
		err := os.WriteFile(m.file, bytes, 0444)
		if err != nil {
			io.Panic(err)
		}
	}
}

func (m *WorkspaceManager) FilterLibs(filter string) []*XsModule {

	var filtered []*XsModule = []*XsModule{}
	for _, modules := range m.workspace.Code {
		for packageName, path := range modules {
			matched, err := PatternMatching(packageName, filter)
			if err != nil {
				io.Println("Error:", err)
				continue
			}

			if matched {
				filtered = append(filtered, &XsModule{Name: packageName, Directory: path})
			}
		}
	}

	return filtered
}

func (m *WorkspaceManager) ExtractModule(name string) *XsModule {
	for _, modules := range m.workspace.Code {
		for packageName, path := range modules {
			if packageName == name {
				return &XsModule{Name: packageName, Directory: path}
			}
		}
	}
	return nil
}

func (m *WorkspaceManager) AddModule(name string, dir string) {
	subDir := strings.Split(dir, "/")[0]
	if m.workspace.Code == nil {
		m.workspace.Code = make(map[string]map[string]string)
	}
	if m.workspace.Code[subDir] == nil {
		m.workspace.Code[subDir] = make(map[string]string)
	}

	m.workspace.Code[subDir][name] = dir

	m.Save()

}

func (m *WorkspaceManager) RemoveModule(name string) {
	for _, modules := range m.workspace.Code {
		for packageName, _ := range modules {
			if packageName == name {
				delete(modules, packageName)
			}
		}
	}
	m.Save()
}

var wsInstance *WorkspaceManager
var wsOnce sync.Once

func GetInstanceWsManager() (*WorkspaceManager, error) {
	wsOnce.Do(func() {
		wsInstance = &WorkspaceManager{}
		wsInstance.file = "." + viper.GetString("files.workspace")
		wsInstance.workspace = &Workspace{}
		err := wsInstance.Load()
		if err != nil {
			io.Panic("Workspace file corrupted: ", wsInstance.file, err)
		}

	})
	return wsInstance, nil
}
