package git

import (
	"path/filepath"
	"xs/pkg/io"
)

const GIT_DIR = ".git"

func CloneGitRepository(url string, path string, bare bool, asModule bool, updateIfExists bool, replaceRemote string) error {
	gitDir := path + "/" + GIT_DIR

	var gt GitInterface

	gt = &GoGit{
		BasePath: path,
		Remote:   url,
	}

	var err error

	if !gt.IsRepoExists() {
		if asModule {
			err = gt.GitAddSubmodule()
		} else {
			var err error
			if bare { // todo need refactoring
				err = gt.GitClone(bare)
			} else {
				err = gt.CloneFromIpfs()
			}
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

				err = gt.SetHead("master")
				if err != nil {
					io.Fatal("Attach head error: " + err.Error())
				}
			}
		}

		if err != nil {
			io.Fatal("Git error: " + err.Error())
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
