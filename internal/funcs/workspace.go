package funcs

import (
	"encoding/json"
	"xs/pkg/io"
	"xs/pkg/tools"
)

type Workspace struct {
	Type     string              `json:"type"`
	Version  int                 `json:"version"`
	Sections map[string]*Section `json:"sections"`
}

type Section struct {
	State    string `json:"state"`
	Template string `json:"template"`
}

type WsManager struct {
	workspace *Workspace
	file      string
}

func (m *WsManager) Load() error {
	fileData, err := tools.ReadFile(m.file)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), m.workspace)
		return nil
	} else {
		return err
	}
}

func (m *WsManager) GetSectionRepository(section string) string {
	return m.workspace.Sections[section].Template
}

func (m *WsManager) SetSectionState(section string, state string) {
	v := m.workspace.Sections[section]
	v.State = state
}

func (m *WsManager) GetSectionState(section string) string {
	return m.workspace.Sections[section].State
}

func (m *WsManager) Save() {
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

func (m *WsManager) GetSections() map[string]*Section {
	return m.workspace.Sections
}

func NewWsManager() (*WsManager, error) {
	manager := WsManager{}
	manager.file = "./xs-workspace.json" //todo move to const
	manager.workspace = &Workspace{}
	err := manager.Load()
	return &manager, err
}
