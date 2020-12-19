package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Swage",
	Long:  `All software has versions. This is Swage's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Swage v%s", swageVersion)
	},
}
