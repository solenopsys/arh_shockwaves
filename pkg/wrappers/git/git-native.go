package git

import (
	"os"
	"os/exec"
	"xs/pkg/io"
)

type NativeGit struct {
	BasePath string
	Remote   string
}

func (g *NativeGit) RunGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	return cmd.Run()
}

func (g *NativeGit) IsRepoExists() bool {
	gitDir := g.BasePath + "/" + GIT_DIR

	_, err := os.Stat(gitDir)

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	io.Println("Error:", err)
	return false
}

func (g *NativeGit) GitClone(bare bool) error {
	return g.RunGitCommand("clone", g.Remote, g.BasePath)
}

func (g *NativeGit) RemoveRemote(name string) error {
	return g.RunGitCommand("remote", "remove", name)
}
func (g *NativeGit) SetRemote(name string, url string) error {
	return g.RunGitCommand("remote", "add", name, url)
}

func (g *NativeGit) GitAddSubmodule() error {
	return g.RunGitCommand("submodule", "add", g.Remote, g.BasePath)
}
func (g *NativeGit) GitUpdateSubmodules() error {
	return g.RunGitCommand("submodule", "update", "--init", "--recursive")
}

func (g *NativeGit) GitUpdate() error {
	return g.RunGitCommand("update", "--init", "--recursive")
}

func (g *NativeGit) UpdateServerInfo() error {
	return g.RunGitCommand("update-server-info")
}

//func (g *NativeGit) UnpackObjects() error
///*
//mv objects/pack/*.pack .
//git unpack-objects < *.pack
//rm -f *.pack objects/pack/*
// */
//	return g.RunGitCommand("unpack-objects")
//}

//func unpack(pack string) {
//
//	git unpack-objects < SAMPLE/*.pack
//
//	}
