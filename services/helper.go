package services

import (
	"encoding/json"
	"os"
)

type PackageTypes string

const (
	Npm PackageTypes = "npm" // Note: The value should be a string literal, not a call to string().
	Git PackageTypes = "git"
)

type Config struct {
	PackageType PackageTypes
}

type InitHelper struct {
	dirs         []string
	PackageType  PackageTypes
	configName   string
	dirsFileName string
}

func (h *InitHelper) createDir(dir string) {
	println("Create dir: ", dir)
	os.Mkdir(dir, 0755)
}

func (h *InitHelper) createDirs() {

	dirs := h.dirs

	for _, dir := range dirs {
		h.createDir(dir)
	}
}
func (h *InitHelper) saveConfigToJson(cofing *Config, file string) {
	configJson, _ := json.Marshal(cofing)
	os.WriteFile(file, configJson, 0644)
}

func (h *InitHelper) createConfig() *Config {
	config := Config{PackageType: h.PackageType}
	h.saveConfigToJson(&config, h.configName)
	return &config
}

func (h *InitHelper) InitRepository() {
	println("Init repository")
	h.loadDirsConfig()
	h.createDirs()
	h.createConfig()
}

func (h *InitHelper) loadDirsConfig() {
	dirs := []string{}
	file, err := os.ReadFile(h.dirsFileName)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &dirs)
	h.dirs = dirs
}

func NewHelper() *InitHelper {
	helper := InitHelper{}
	helper.PackageType = Git
	helper.configName = "config/xs.json"
	helper.dirsFileName = "config/folders.json"
	return &helper
}
