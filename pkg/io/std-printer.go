package io

import (
	"os/exec"
)

type StdPrinter struct {
	Key            string
	Command        string
	Args           []string
	PrintToConsole bool
}

func (s *StdPrinter) Start() int {
	store := GetLogStore()
	cmd := exec.Command(s.Command, s.Args...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		Panic(err)
	}
	if err = cmd.Start(); err != nil {
		Panic(err)
	}
	for {
		tmp := make([]byte, 2048)
		n, err := stdout.Read(tmp)
		res := string(tmp[:n])

		store.MessagesStream <- LogMessage{Message: res, Key: s.Key}
		if err != nil {
			break
		}
	}

	cmd.Wait()
	resultCode := cmd.ProcessState.ExitCode()
	return resultCode
}
