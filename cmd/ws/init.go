package ws

import (
	"github.com/spf13/cobra"
	"xs/internal/configs"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "Workspace initialization",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		configs.LoadWorkspace("https://github.com/solenopsys/tp-workspace.git")
	},
}
