package cmd

import (
	"errors"
	"fmt"

	"github.com/markruler/swage/pkg/excel"
	"github.com/markruler/swage/pkg/parser"
	"github.com/spf13/cobra"
)

var (
	outputPath string
	template   string
)

var genCmd = &cobra.Command{
	Use:   "gen [JSON_PATH]",
	Short: "Generate a Excel file",
	Long: `Generate a Excel file
		
ex) swage gen aio/example/example.json -o $HOME/swage.xlsx
`,
	RunE: genRun,
}

func init() {
	genCmd.Flags().StringVarP(&outputPath, "output", "o", "swage.xlsx", "set a path to save a Excel file")
	genCmd.Flags().StringVarP(&template, "template", "t", "1", "set a Excel template [1]")
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

	p := parser.New(args[0])
	var err error

	swaggerAPI, err := p.Parse()
	if err != nil {
		return err
	}

	xl := excel.New()

	if err = xl.Generate(swaggerAPI, template); err != nil {
		return err
	}

	if err := xl.File.SaveAs(outputPath); err != nil {
		return err
	}

	if verbose {
		fmt.Printf("OUTPUT >>> %s\n", outputPath)
	}
	return nil
}
