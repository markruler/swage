package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Update this whenever making a new release.
	// (with VERSION file)
	swageVersion = "0.1.0"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Swage",
	Long:  `All software has versions. This is Swage's`,
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("swage v%s\n", swageVersion)
}
