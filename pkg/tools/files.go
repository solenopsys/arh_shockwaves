package tools

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"xs/pkg/io"
)

func isWorkspaceRootDir(dir string) bool {
	return Exists(dir + viper.GetString("files.workspace")) // todo move to constant
}

func CheckWorkspace(dir string, before string) (bool, string) {
	isWorkspaceRootDir := isWorkspaceRootDir(dir)
	if before == dir {
		return false, dir
	}
	if isWorkspaceRootDir {
		return true, dir
	} else {
		parentDir := filepath.Dir(dir)
		return CheckWorkspace(parentDir, dir)
	}
}

func FindWorkspaceRootDir() (bool, string) {
	dir, err := os.Getwd()
	if err != nil {
		io.Fatal(err)
	}

	return CheckWorkspace(dir, "")
}

func ToWorkspaceRootDir() error {
	isWorkspaceRootDir, dir := FindWorkspaceRootDir()
	if !isWorkspaceRootDir {
		return errors.New("Not in workspace root dir: " + dir)
	}

	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	return nil
}

func CreateDirs(dir string) error {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		io.Println(err)
		return err
	}

	return nil
}

func Exists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return true

}

func DeleteDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		io.Println(err)
		return err
	}
	return err
}

func ClearDirectory(dir string) error {

	err := DeleteDir(dir)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		io.Println(err)
		return err
	}

	return nil
}
