package build

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
	"xs/services"
)

type StdPrinter struct {
	out     chan string
	command string
	args    []string
}

func (s *StdPrinter) processing() {

	for s.out != nil {
		select {
		case res := <-s.out:
			r := strings.Replace(res, "\n", "\r\n", -1)
			fmt.Print(r)
		}
	}
}

func (s *StdPrinter) start() {
	cmd := exec.Command(s.command, s.args...)
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

		s.out <- res
		if err != nil {
			break
		}
	}
}

var cmdContainer = &cobra.Command{
	Use:   "container [name]",
	Short: "Containers for module build and push to registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		m := args[0]
		groupDir := "modules"

		mod, extractError := services.ExtractModule(m, groupDir)
		if extractError != nil {
			println("Error", extractError.Error())
			return
		}
		path := "./" + groupDir + "/" + mod.Directory

		platform := "amd64"
		reg := "registry.solenopsys.org"

		errDir := os.Chdir(path)
		if errDir != nil {
			fmt.Println(errDir)
			return
		}

		command := "nerdctl"
		println("command:" + command)

		arg := "build --platform=" + platform + " --output type=image,name=" + reg + "/" + mod.Directory + ":latest,push=true ."
		argsSplit := strings.Split(arg, " ")
		if errDir != nil {
			fmt.Println(errDir)
			return
		}
		var changeDir chan string = make(chan string)
		var stdPrinter StdPrinter = StdPrinter{out: changeDir, command: "nerdctl", args: argsSplit}
		go stdPrinter.processing()
		stdPrinter.start()
	},
}
