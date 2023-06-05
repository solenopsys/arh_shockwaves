package utils

import "encoding/json"

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

func (m *WsManager) Load() {
	fileData, err := ReadFile(m.file)
	if err == nil {
		err = json.Unmarshal([]byte(fileData), m.workspace)
	}
	if err != nil {
		panic(err)
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
		panic(err)
	} else {
		err := WriteFile(m.file, bytes)
		if err != nil {
			panic(err)
		}
	}
}

func (m *WsManager) GetSections() map[string]*Section {
	return m.workspace.Sections
}

func NewWsManager() *WsManager {
	manager := WsManager{}
	manager.file = "./workspace.json"
	manager.workspace = &Workspace{}
	manager.Load()
	return &manager
}
