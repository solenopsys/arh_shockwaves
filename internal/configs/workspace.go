package configs

import (
	"encoding/json"
	"xs/pkg/io"
	"xs/pkg/tools"
)

type Format struct {
	Type    string `json:"type"`
	Version int    `json:"version"`
}

type Workspace struct {
	Format    Format
	Templates map[string]map[string]string `json:"templates"`
	Code      map[string]map[string]string `json:"code"`
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

func (m *WsManager) GetTemplateDirectory(dir string) string {
	return m.workspace.Templates["sections"][dir]
}

func NewWsManager() (*WsManager, error) {
	manager := WsManager{}
	manager.file = "./xs-workspace.json" //todo move to const
	manager.workspace = &Workspace{}
	err := manager.Load()
	if err != nil {
		io.Panic("Workspace file corrupted: ", manager.file, err)
	}
	return &manager, err
}
