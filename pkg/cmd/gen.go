package cmd

import (
	"errors"
	"fmt"

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
	RunE: genRun,
}

func init() {
	genCmd.Flags().StringVarP(&outputPath, "output", "o", "", "set a path to save a Excel file")
	genCmd.Flags().BoolP("verbose", "v", false, "verbose print")
}

func genRun(cmd *cobra.Command, args []string) error {
	verbose, _ := cmd.Flags().GetBool("verbose")
	if len(args) == 0 {
		return errors.New("JSON_PATH is required")
	}
	if verbose {
		fmt.Printf(">>> INPUT %s\n", args[0])
	}
	swaggerAPI, err := parser.Parse(args[0])
	if err != nil {
		return err
	}
	path, err := excel.Save(swaggerAPI, outputPath, verbose)
	if err != nil {
		return err
	}
	fmt.Printf("OUTPUT >>> %s\n", path)
	return nil
}
