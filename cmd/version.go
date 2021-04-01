package cmd

import (
	"github.com/spf13/cobra"
)

var swageVersion = "SNAPSHOT"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Swage version information",
	Long:  `Show the Swage version information`,
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	cmd.Printf("swage %s\n", swageVersion)
}
