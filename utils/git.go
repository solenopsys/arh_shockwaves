package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const GIT_DIR = ".git"

func CloneGitRepository(url string, path string, asModule bool, updateIfExists bool) error {
	gitDir := path + "/" + GIT_DIR

	gt := &GitTools{
		basePath: path,
		remote:   url,
	}

	if !gt.IsRepoExists() {
		println("Clone repository: " + url)
		if asModule {
			return gt.GitAddSubmodule()
		} else {
			return gt.GitClone()
		}
	} else {
		println("Repo exists: " + url)
		if updateIfExists {
			if asModule {
				fullPath, _ := filepath.Abs(gitDir)
				println("Repository already exists update: " + fullPath)
				return gt.GitUpdateSubmodules()
			} else {
				return gt.GitUpdate()
			}
		} else {
			return nil
		}
	}

}

type GitTools struct {
	basePath string
	remote   string
}

func (g *GitTools) RunGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	return cmd.Run()
}

func (g *GitTools) IsRepoExists() bool {
	gitDir := g.basePath + "/" + GIT_DIR
	_, err := syscall.Open(gitDir, syscall.O_RDONLY, 0)

	if err != nil {
		panic(err.Error())
	}
	return os.IsNotExist(err)
}

func (g *GitTools) GitClone() error {
	return g.RunGitCommand("clone", g.remote, g.basePath)
}

func (g *GitTools) GitAddSubmodule() error {
	return g.RunGitCommand("submodule", "add", g.remote, g.basePath)
}
func (g *GitTools) GitUpdateSubmodules() error {
	return g.RunGitCommand("submodule", "update", "--init", "--recursive")
}

func (g *GitTools) GitUpdate() error {
	return g.RunGitCommand("update", "--init", "--recursive")
}
