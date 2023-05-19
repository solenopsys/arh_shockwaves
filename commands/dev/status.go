package dev

import (
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Show status of installed env programs (git,nx,npm,go,...)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		commands := map[string][]string{
			"git":     {"git", "version"},
			"npm":     {"npm", "-v"},
			"go":      {"go", "version"},
			"nx":      {"npm", "show @nrwl/cli version"},
			"nerdctl": {"nerdctl", "version"},
		}
		for name, command := range commands {
			arg := command[1]
			splitArg := strings.Split(arg, " ")
			var version, err = exec.Command(command[0], splitArg...).Output()
			if err == nil {
				verLine := string(version)
				//replace ver
				verLine = strings.Replace(verLine, "version", "", 1)
				verLine = strings.Replace(verLine, name, "", 1)
				// trim
				verLine = strings.TrimSpace(verLine)
				println("")
				println(name)
				println(" -------------------------------->")
				println(verLine)
			} else {
				println(name+":", "not installed")
			}

		}

	},
}
