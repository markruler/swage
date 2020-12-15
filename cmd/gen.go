package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swage",
	Short: "Swage is a swagger.json converter to excel format",
	Long: `Swage is a swagger.json converter
								to excel format`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("%s\n", "first argument is required as a 'json path'")
		}
		swaggerAPI, err := parse(args[0])
		if swaggerAPI == nil || err != nil {
			log.Fatal(err)
		}
	},
}

// Execute ...
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
