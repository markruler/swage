package cmd

import (
	"github.com/spf13/cobra"
)

const swageVersion = "0.1.0"

var swageCmd = &cobra.Command{
	Use:   "swage",
	Short: "Swage is a swagger.json converter to excel format",
	Long: `Swage is a swagger.json converter
(to Excel format)

ex) swage gen aio/example/example.json -o $HOME/swage.xlsx
`,
	SilenceUsage: true,
}

func init() {
	swageCmd.AddCommand(genCmd)
	swageCmd.AddCommand(versionCmd)
}

// Execute swage
func Execute() {
	swageCmd.Execute()
}
