package cmd

import (
	"errors"
	"strings"

	"github.com/markruler/swage/converter"
	"github.com/markruler/swage/spec"
	"github.com/markruler/swage/xlsx"
	"github.com/markruler/swage/xlsx/simple"
	"github.com/spf13/cobra"
)

var (
	outputPath   string
	templateName string
)

var genCmd = &cobra.Command{
	Use:   "gen PATH",
	Short: "Generate an XLSX file locally",
	Long: `
Generate an XLSX file locally
`,
	RunE: genRun,
}

func init() {
	genCmd.Flags().StringVarP(&outputPath, "output", "o", "swage.xlsx", "set a path to save a Excel file")
	genCmd.Flags().StringVarP(&templateName, "template", "t", xlsx.Simple, "set a Excel template [simple]")
	genCmd.Flags().BoolP("verbose", "v", false, "verbose print")
}

func genRun(cmd *cobra.Command, args []string) error {
	verbose, _ := cmd.Flags().GetBool("verbose")
	if len(args) == 0 {
		return errors.New("PATH is required")
	}
	sourcePath := args[0]
	if verbose {
		cmd.Printf(">>> INPUT %s\n", sourcePath)
	}

	swaggerAPI, err := spec.Parse(sourcePath)
	if err != nil {
		return err
	}
	// TEST
	converter.Convert(swaggerAPI)

	var tmpl xlsx.Template

	switch strings.TrimSpace(templateName) {
	case xlsx.Simple:
		tmpl = simple.New()
	// TODO:
	// case xlsx.Print:
	// 	tmpl = print.New()
	default:
		return errors.New("the template not found")
	}

	if err = tmpl.Generate(swaggerAPI); err != nil {
		return err
	}

	if err := tmpl.GetExcel().File.SaveAs(outputPath); err != nil {
		return err
	}

	if verbose {
		cmd.Printf("OUTPUT >>> %s\n", outputPath)
	}
	return nil
}
