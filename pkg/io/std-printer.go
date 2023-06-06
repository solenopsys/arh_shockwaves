package io

import (
	"os/exec"
	"strings"
)

type StdPrinter struct {
	Out            chan string
	Command        string
	Args           []string
	PrintToConsole bool
}

func (s *StdPrinter) Processing() {

	for s.Out != nil {
		select {
		case res := <-s.Out:
			r := strings.Replace(res, "\n", "\r\n", -1)
			if s.PrintToConsole {
				Print(r)
			}
		}
	}
}

func (s *StdPrinter) Start() int {
	cmd := exec.Command(s.Command, s.Args...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		panic(err)
	}
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	for {
		tmp := make([]byte, 2048)
		n, err := stdout.Read(tmp)
		res := string(tmp[:n])
		//replace multiple spaces with one

		s.Out <- res
		if err != nil {
			break
		}
	}

	cmd.Wait()
	resultCode := cmd.ProcessState.ExitCode()
	return resultCode
}
