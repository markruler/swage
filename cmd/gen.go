package cmd

import (
	"errors"
	"strings"

	"github.com/markruler/swage/parser"
	"github.com/markruler/swage/template"
	"github.com/markruler/swage/template/simple"
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
	genCmd.Flags().StringVarP(&templateName, "template", "t", template.Simple, "set a Excel template [simple]")
	genCmd.Flags().BoolP("verbose", "v", false, "verbose print")
}

func genRun(cmd *cobra.Command, args []string) error {
	verbose, _ := cmd.Flags().GetBool("verbose")
	if len(args) == 0 {
		return errors.New("PATH is required")
	}
	if verbose {
		cmd.Printf(">>> INPUT %s\n", args[0])
	}

	p := parser.New(args[0])
	var err error

	swaggerAPI, err := p.Parse()
	if err != nil {
		return err
	}

	var tmpl template.Template

	switch strings.TrimSpace(templateName) {
	case template.Simple:
		tmpl = simple.New()
	// TODO:
	// case "print":
	// 	template = print.New()
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
