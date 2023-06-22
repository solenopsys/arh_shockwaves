package serve

import (
	"github.com/spf13/cobra"
	"xs/pkg/io"
)

var cmdFront = &cobra.Command{
	Use:   "front [name]",
	Short: "Frontend serve",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		io.StartProxy("C:\\dev\\sources\\MAIN\\temp2\\frontends\\dist\\fronts\\fr-web")

	},
}
