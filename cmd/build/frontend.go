package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/internal/extractors"
	"xs/internal/services"
)

var cmdFrontend = &cobra.Command{
	Use:   "frontend [name]",
	Short: "frontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		executor := compilers.Frontend{PrintConsole: true}
		extractor := extractors.Frontend{}
		services.CompileGroup(args[0], "entrances", executor, extractor)
	},
}
