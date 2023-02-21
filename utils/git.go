package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func CloneGitRepository(url string, path string, asModule bool) error {
	gitDir := path + "/.git"

	println("Clone repository: " + url + " to " + path)
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		if asModule {
			cmd := exec.Command("git", "submodule", "add", url, path)
			return cmd.Run()
		} else {
			cmd := exec.Command("git", "clone", url, path)
			return cmd.Run()
		}

	} else {
		fullPath, _ := filepath.Abs(gitDir)
		println("Repository already exists update: " + fullPath)
		cmd := exec.Command("git", "pull")
		return cmd.Run()
	}

}
