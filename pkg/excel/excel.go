package excel

import (
	"fmt"
	"log"

	"github.com/markruler/swage/pkg/spec"
)

const (
	defaultFilePath = "swage.xlsx"
)

// Save ...
func Save(swaggerAPI *spec.SwaggerAPI, outputFilePath string, verbose bool) {
	if swaggerAPI == nil {
		return
	}
	xl := createIndexSheet(swaggerAPI)
	// fmt.Println(xl)
	// fmt.Println(xl.GetDocProps())
	var path string
	if outputFilePath != "" {
		path = outputFilePath
	} else {
		path = defaultFilePath
	}

	if err := xl.SaveAs(path); err != nil {
		log.Fatalln(err)
	}
	if verbose {
		fmt.Printf("OUTPUT >>> %s\n", path)
	}
}
