package excel

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/markruler/swage/pkg/spec"
)

var (
	excelFileName = "swage.xlsx"
)

// Save ...
func Save(swaggerAPI *spec.SwaggerAPI, outputPath string, verbose bool) {
	if swaggerAPI == nil {
		return
	}
	xl := createIndexSheet(swaggerAPI)

	if outputPath == "" {
		setOutputPath(xl, excelFileName, verbose)
	} else {
		setOutputPath(xl, outputPath, verbose)
	}
}

func setOutputPath(xl *excelize.File, path string, verbose bool) {
	if err := xl.SaveAs(path); err != nil {
		log.Fatalln(err)
	}
	if verbose {
		fmt.Printf("OUTPUT >>> %s\n", path)
	}
}
