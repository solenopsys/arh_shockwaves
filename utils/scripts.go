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

func readFile(url string) ([]byte, error) {
	file, err := os.Open(url)
	if err != nil {
		return nil, err
	} else {
		return ioutil.ReadAll(file)
	}
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
	httpBody, err := downloadScript(file)
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
