package serve

import (
	"github.com/spf13/cobra"
	"strconv"
	"sync"
	"xs/pkg/io"
)

var cmdFront = &cobra.Command{
	Use:   "front [name]",
	Short: "Frontend serve",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		count := len(args)

		port := 8000

		wg := new(sync.WaitGroup)
		for i := 0; i < count; i++ {

			name := args[i]

			strPort := strconv.Itoa(port)

			io.Println("Print: " + name + ": " + strPort)
			go io.StartProxy(".\\dist\\fronts\\"+name, strPort)
			port++
			wg.Add(1)
		}

		wg.Wait()
	},
}
