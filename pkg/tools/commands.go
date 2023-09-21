package tools

import (
	"bytes"
	sio "io"
	"xs/pkg/io"

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

func CommandApplyFromUrl(url string, command string) {

	httpBody, err := downloadScript(url)
	if err != nil {
		io.Println(err.Error())
		return
	}

	execCommand(command, httpBody)
}

func CommandApplyFromFile(file string, command string) {
	io.Println("Start install")
	httpBody, err := os.ReadFile(file)
	if err != nil {
		io.Println(err.Error())
		return
	}

	execCommand(command, httpBody)
}

func execCommand(command string, httpBody []byte) {
	cmdIn := exec.Command(command)
	cmdIn.Stdin = bytes.NewBuffer(httpBody)

	stdout, err := cmdIn.StdoutPipe()
	if err != nil {
		io.Println(err)
	}

	// Start the command
	err = cmdIn.Start()
	if err != nil {
		io.Println(err)
	}

	// Use io.Copy to print the command's output in real-time
	_, err = sio.Copy(os.Stdout, stdout)
	if err != nil {
		io.Println(err)
	}

	// Wait for the command to finish
	err = cmdIn.Wait()
	if err != nil {
		io.Println(err)
	}
}
