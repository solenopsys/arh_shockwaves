package wrappers

import (
	"path/filepath"
	"xs/pkg/io"
	"xs/pkg/wrappers/git"
)

const GIT_DIR = ".git"

func CloneGitRepository(url string, path string, asModule bool, updateIfExists bool, replaceRemote string) error {
	gitDir := path + "/" + GIT_DIR

	var gt git.GitInterface

	gt = &git.GoGit{
		BasePath: path,
		Remote:   url,
	}

	var err error

	if !gt.IsRepoExists() {
		if asModule {
			err = gt.GitAddSubmodule()
		} else {

			err = gt.GitClone()
			if err != nil {
				io.Fatal("Git clone error: " + err.Error())
			}
			if replaceRemote != "" {
				err = gt.RemoveRemote("origin")
				if err != nil {
					io.Fatal("Git origin remove error: " + err.Error())
				}
				err = gt.SetRemote("origin", replaceRemote)
				if err != nil {
					io.Fatal("Git origin add error: " + err.Error())
				}
			}
		}

		if err != nil {
			io.Fatal("Git error: " + err.Error())
		} else {
			io.Println("Cloned repository: " + url)
		}

		return err
	} else {

		if updateIfExists {
			if asModule {
				fullPath, _ := filepath.Abs(gitDir)
				err = gt.GitUpdateSubmodules()
				io.Println("Exists repo updated: " + fullPath)
			} else {
				err = gt.GitUpdate()
				io.Println("Exists repo updated: " + url)
			}
		}

		return err

	}

}
