package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var swageVersion = "SNAPSHOT"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Swage",
	Long:  `All software has versions. This is Swage's`,
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("swage %s\n", swageVersion)
}
