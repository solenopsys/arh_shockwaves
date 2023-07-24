package tools

import (
	"os"
	"xs/pkg/io"
)

type PathTools struct {
	basePath string
}

func (p *PathTools) SetBasePathPwd() {
	currentDir, errDir := os.Getwd()
	if errDir != nil {
		io.Fatal(errDir)
	}
	p.basePath = currentDir
}

func (p *PathTools) MoveTo(path string) {
	errDir := os.Chdir(path)
	if errDir != nil {
		io.Fatal(errDir, path)
	}
}

func (p *PathTools) MoveToBasePath() {
	errDir := os.Chdir(p.basePath)
	if errDir != nil {
		io.Fatal(errDir)
	}
}

func (p *PathTools) GetPwd() string {
	currentDir, _ := os.Getwd()

	return currentDir
}

func NewPathTools() *PathTools {
	p := &PathTools{}
	p.SetBasePathPwd()
	return p
}
