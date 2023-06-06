package tools

import "os"

type PathTools struct {
	basePath string
}

func (p *PathTools) SetBasePathPwd() {
	currentDir, errDir := os.Getwd()
	if errDir != nil {
		panic(errDir)
	}
	p.basePath = currentDir
}

func (p *PathTools) MoveTo(path string) {
	errDir := os.Chdir(path)
	if errDir != nil {
		panic(errDir)
	}
}

func (p *PathTools) MoveToBasePath() {
	errDir := os.Chdir(p.basePath)
	if errDir != nil {
		panic(errDir)
	}
}

func NewPathTools() *PathTools {
	p := &PathTools{}
	p.SetBasePathPwd()
	return p
}
