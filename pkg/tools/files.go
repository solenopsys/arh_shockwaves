package tools

import (
	"os"
	"xs/pkg/io"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return []byte(data), err
}

func WriteFile(fileName string, data []byte) error {
	err := os.WriteFile(fileName, data, 0444)
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

func DirExists(dir string) bool {
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
