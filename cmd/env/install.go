package env

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
)

var cmdInstall = &cobra.Command{
	Use:   "install",
	Short: "Install all necessary programs (git,nx,npm,go,...)",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		io.Println("NOT IMPLEMENTED")
	},
}
