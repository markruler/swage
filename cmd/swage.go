package cmd

import (
	"github.com/spf13/cobra"
)

var swageCmd = &cobra.Command{
	Use:   "swage",
	Short: "Swage convert OpenAPI Schema to XLSX",
	Long: `
Swage convert OpenAPI Schema to XLSX
`,
	SilenceUsage: false,
}

func init() {
	swageCmd.AddCommand(genCmd)
	swageCmd.AddCommand(versionCmd)
}

// Execute swage
func Execute() error {
	return swageCmd.Execute()
}
