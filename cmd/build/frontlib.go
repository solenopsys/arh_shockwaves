package build

import (
	"github.com/spf13/cobra"
	"xs/internal/compilers"
	"xs/internal/extractors"
	"xs/internal/services"
)

var cmdFrontlib = &cobra.Command{
	Use:   "frontlib [name]",
	Short: "Frontend build",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		executor := compilers.Frontlib{PrintConsole: true}
		extractor := extractors.Frontlib{}
		services.CompileGroup(args[0], "packages", executor, extractor)
	},
}
