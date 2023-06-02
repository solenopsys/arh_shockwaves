package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func downloadScript(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		return ioutil.ReadAll(response.Body)
	}
}

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
		fmt.Println(err)
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
		fmt.Println(err)
		return err
	}
	return err
}

func ClearDirectory(dir string) error {

	err := DeleteDir(dir)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func CommandApplyFromUrl(url string, command string) {

	httpBody, err := downloadScript(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	execCommand(command, httpBody)
}

func CommandApplyFromFile(file string, command string) {
	fmt.Println("Start install")
	httpBody, err := ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	execCommand(command, httpBody)
}

func execCommand(command string, httpBody []byte) {
	cmdIn := exec.Command(command)
	cmdIn.Stdin = bytes.NewBuffer(httpBody)

	stdout, err := cmdIn.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	// Start the command
	err = cmdIn.Start()
	if err != nil {
		fmt.Println(err)
	}

	// Use io.Copy to print the command's output in real-time
	_, err = io.Copy(os.Stdout, stdout)
	if err != nil {
		fmt.Println(err)
	}

	// Wait for the command to finish
	err = cmdIn.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
