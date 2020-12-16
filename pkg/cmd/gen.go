package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/markruler/swage/pkg/excel"
	"github.com/markruler/swage/pkg/parser"
	"github.com/spf13/cobra"
)

const swageVersion = "0.1.0"

var outputPath string

var swageCmd = &cobra.Command{
	Use:   "swage",
	Short: "Swage is a swagger.json converter to excel format",
	Long: `Swage is a swagger.json converter
(to Excel format)

ex) swage gen aio/example/example.json -o $HOME/swage.xlsx
`,
	SilenceUsage: true,
}

var genCmd = &cobra.Command{
	Use:   "gen [JSON_PATH]",
	Short: "Generate a Excel file",
	Long: `Generate a Excel file
		
ex) swage gen aio/example/example.json -o $HOME/swage.xlsx
`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if len(args) == 0 {
			log.Fatalf("%s\n", "JSON_PATH is required")
		}
		if verbose {
			fmt.Printf(">>> INPUT %s\n", args[0])
		}
		swaggerAPI, err := parser.Parse(args[0])
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		excel.Save(swaggerAPI, outputPath, verbose)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Swage",
	Long:  `All software has versions. This is Swage's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Swage v%s", swageVersion)
	},
}

func init() {
	swageCmd.AddCommand(genCmd)
	swageCmd.AddCommand(versionCmd)
	genCmd.Flags().StringVarP(&outputPath, "output", "o", "", "set a path to save a Excel file")
	genCmd.Flags().BoolP("verbose", "v", false, "verbose print")
}

// Execute ...
func Execute() error {
	if err := swageCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
