package utils

import "os/exec"

func CloneGitRepository(url string, path string) error {
	cmd := exec.Command("git", "clone", url, path)
	return cmd.Run()
}
