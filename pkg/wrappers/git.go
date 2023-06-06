package wrappers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const GIT_DIR = ".git"

func CloneGitRepository(url string, path string, asModule bool, updateIfExists bool) error {
	gitDir := path + "/" + GIT_DIR

	gt := &GitTools{
		basePath: path,
		remote:   url,
	}

	var err error

	if !gt.IsRepoExists() {
		if asModule {
			err = gt.GitAddSubmodule()
		} else {
			err = gt.GitClone()
		}
		println("Cloned repository: " + url)

		return err
	} else {

		if updateIfExists {
			if asModule {
				fullPath, _ := filepath.Abs(gitDir)
				err = gt.GitUpdateSubmodules()
				println("Exists repo updated: " + fullPath)
			} else {
				err = gt.GitUpdate()
				println("Exists repo updated: " + url)
			}
		}

		return err

	}

}

type GitTools struct {
	basePath string
	remote   string
}

func (g *GitTools) RunGitCommand(args ...string) error {
	println("RUN GIT git " + strings.Join(args, " "))
	cmd := exec.Command("git", args...)
	return cmd.Run()
}

func (g *GitTools) IsRepoExists() bool {
	gitDir := g.basePath + "/" + GIT_DIR

	_, err := os.Stat(gitDir)

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	fmt.Println("Error:", err)
	return false
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