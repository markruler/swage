package excel

import (
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/markruler/swage/pkg/spec"
	"github.com/markruler/swage/pkg/style"
)

var (
	indexSheetName = "INDEX"
	excelFileName  = "swage.xlsx"
)

// Save ...
func Save(swaggerAPI *spec.SwaggerAPI, outputPath string, verbose bool) {
	if swaggerAPI == nil {
		return
	}
	if verbose {
		log.Printf("[excel.go] swaggerAPI: %v\n", swaggerAPI)
	}
	xl := excelize.NewFile()

	/////////////////
	// INDEX Sheet //
	/////////////////
	xl.SetSheetName("Sheet1", indexSheetName)

	// index := xl.GetSheetIndex(indexSheetName)
	// xl.SetActiveSheet(index)

	// Index Sheet Column Style
	xl.SetColStyle(indexSheetName, "A", style.Center(xl))
	xl.SetColStyle(indexSheetName, "B", style.Center(xl))
	xl.SetColStyle(indexSheetName, "C", style.Center(xl))
	xl.SetColStyle(indexSheetName, "D", style.Left(xl))
	xl.SetColStyle(indexSheetName, "E", style.Left(xl))
	xl.SetColWidth(indexSheetName, "D", "E", 45.0)

	// Index Sheet Header
	xl.SetCellStr(indexSheetName, "A1", "#")
	xl.SetCellStr(indexSheetName, "B1", "tag")
	xl.SetCellStr(indexSheetName, "C1", "method")
	xl.SetCellStr(indexSheetName, "D1", "path")
	xl.SetCellStr(indexSheetName, "E1", "summary")
	
	// Index Sheet Data
	var row int
	for path, spec := range swaggerAPI.Paths {
		row++
		// fmt.Printf("%s:\n%v\n", path, spec)
		// TODO:
		xl.SetCellInt(indexSheetName, fmt.Sprintf("%s%d", "A", row+1), row)
		xl.SetCellHyperLink(indexSheetName, fmt.Sprintf("%s%d", "A", row+1), fmt.Sprintf("%d!A1", row), "Location")
		xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "B", row+1), spec.Get.Tags[0])
		xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "C", row+1), "GET")
		xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "D", row+1), path)
		xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "E", row+1), spec.Get.Summary)
		xl.NewSheet(strconv.Itoa(row))
	}
	
	/////////////////////
	// TODO: API Sheet //
	/////////////////////

	////////////
	// OUTPUT //
	////////////
	if outputPath == "" {
		if err := xl.SaveAs(excelFileName); err != nil {
			log.Fatalln(err)
		}
		if verbose {
			fmt.Printf("OUTPUT >>> %s\n", excelFileName)
		}
	} else {
		if err := xl.SaveAs(outputPath); err != nil {
			log.Fatalln(err)
		}
		if verbose {
			fmt.Printf("OUTPUT >>> %s\n", outputPath)
		}
	}
}
