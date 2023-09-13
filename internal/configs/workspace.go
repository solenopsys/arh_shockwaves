package configs

import (
	"encoding/json"
	"regexp"
	"strings"
	"xs/pkg/io"
	"xs/pkg/tools"
)

type Format struct {
	Type    string `json:"type"`
	Version int    `json:"version"`
}

type XsModule struct {
	Directory string
	Name      string
}

type Workspace struct {
	Format    Format
	Templates map[string]map[string]string `json:"templates"`
	Code      map[string]map[string]string `json:"code"`
}

type WorkspaceManager struct {
	workspace *Workspace
	file      string
}

func (m *WorkspaceManager) Load() error {
	fileData, err := tools.ReadFile(m.file)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), m.workspace)
		return nil
	} else {
		return err
	}
}

func (m *WorkspaceManager) Save() {
	bytes, err := json.MarshalIndent(m.workspace, "", "  ")
	if err != nil {
		io.Panic(err)
	} else {
		err := tools.WriteFile(m.file, bytes)
		if err != nil {
			io.Panic(err)
		}
	}
}

func (m *WorkspaceManager) GetTemplateDirectory(dir string) string {
	return m.workspace.Templates["sections"][dir]
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

func NewWsManager() (*WorkspaceManager, error) {
	manager := WorkspaceManager{}
	manager.file = "./xs-workspace.json" //todo move to const
	manager.workspace = &Workspace{}
	err := manager.Load()
	if err != nil {
		io.Panic("Workspace file corrupted: ", manager.file, err)
	}
	return &manager, err
}
