package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func CloneGitRepository(url string, path string, asModule bool, updateIfExists bool) error {
	gitDir := path + "/.git"

	if _, err := syscall.Open(gitDir, syscall.O_RDONLY, 0); os.IsNotExist(err) {
		println("Clone repository: " + url + " to " + path)
		if asModule {
			cmd := exec.Command("git", "submodule", "add", url, path)
			return cmd.Run()
		} else {
			cmd := exec.Command("git", "clone", url, path)
			return cmd.Run()
		}
	} else {
		println("Repo exists: " + url + " to " + path)
		if updateIfExists {
			if asModule {
				fullPath, _ := filepath.Abs(gitDir)
				println("Repository already exists update: " + fullPath)
				cmd := exec.Command("git", "submodule", "update", "--init", "--recursive")
				return cmd.Run()
			} else {
				cmd := exec.Command("git", "update", "--init", "--recursive")
				return cmd.Run()
			}
		} else {
			return nil
		}
	}

}
