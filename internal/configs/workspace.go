package configs

import (
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
	"sync"
	"xs/pkg/io"
	"xs/pkg/tools"
)

type XsModule struct {
	Directory string
	Name      string
}

type Workspace struct {
	Format string                       `json:"format"`
	Code   map[string]map[string]string `json:"code"`
}

type WorkspaceManager struct {
	workspace *Workspace
	file      string
}

func (m *WorkspaceManager) Load() error {
	fileData, err := tools.ReadFile(m.file)
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
		err := tools.WriteFile(m.file, bytes)
		if err != nil {
			io.Panic(err)
		}
	}
}

func (m *WorkspaceManager) FilterLibs(filter string) []*XsModule {

	var filtered []*XsModule = []*XsModule{}
	for _, modules := range m.workspace.Code {
		for packageName, path := range modules {

			pattern := strings.Replace(filter, "*", ".*", -1)
			matched, err := regexp.MatchString("^"+pattern+"$", packageName)
			if err != nil {
				io.Println("Error:", err)
				continue
			}

			if matched {
				filtered = append(filtered, &XsModule{Name: packageName, Directory: path})
			}
		}
	}

	io.Println("Found  lib count:", len(filtered))
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

var wsInstance *WorkspaceManager
var wsOnce sync.Once

func GetInstanceWsManager() (*WorkspaceManager, error) {
	wsOnce.Do(func() {
		wsInstance = &WorkspaceManager{}
		wsInstance.file = "./workspace.yaml" //todo move to const
		wsInstance.workspace = &Workspace{}
		err := wsInstance.Load()
		if err != nil {
			io.Panic("Workspace file corrupted: ", wsInstance.file, err)
		}

	})
	return wsInstance, nil
}
