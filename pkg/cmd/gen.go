package cmd

import (
	"fmt"
	"log"

	"github.com/markruler/swage/pkg/excel"
	"github.com/markruler/swage/pkg/parser"
	"github.com/spf13/cobra"
)

var outputPath string

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

func init() {
	genCmd.Flags().StringVarP(&outputPath, "output", "o", "", "set a path to save a Excel file")
	genCmd.Flags().BoolP("verbose", "v", false, "verbose print")
}
