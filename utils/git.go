package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func CloneGitRepository(url string, path string) error {
	gitDir := path + "/.git"

	println("Clone repository: " + url + " to " + path)
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", url, path)
		return cmd.Run()
	} else {
		fullPath, _ := filepath.Abs(gitDir)
		println("Repository already exists update: " + fullPath)
		cmd := exec.Command("git", "pull")
		return cmd.Run()
	}

}
