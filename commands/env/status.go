package env

import (
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Show status of installed env programs (git,pnpm,go,...)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		commands := map[string][]string{
			"git":        {"git", "version"},
			"pnpm":       {"pnpm", "-v"},
			"go":         {"go", "version"},
			"ng-packagr": {"ng-packagr", "-v"},
			"nerdctl":    {"nerdctl", "version"},
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
