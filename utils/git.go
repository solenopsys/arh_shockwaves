package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func CloneGitRepository(url string, path string) error {
	gitDir := path + "/.git"

	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", url, path)
		return cmd.Run()
	} else {
		fullpath, _ := filepath.Abs(gitDir)
		println("Repository already exists: " + fullpath)
		return nil
	}

}
