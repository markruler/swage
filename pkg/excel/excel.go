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
	xl := createIndexSheet(swaggerAPI)

	if outputPath == "" {
		setOutputPath(xl, excelFileName, verbose)
	} else {
		setOutputPath(xl, outputPath, verbose)
	}
}

func createIndexSheet(swaggerAPI *spec.SwaggerAPI) *excelize.File {
	xl := excelize.NewFile()
	xl.SetSheetName("Sheet1", indexSheetName)
	xl.SetPanes(indexSheetName, `{
    "freeze": true,
    "split": false,
    "x_split": 1,
    "y_split": 1,
    "top_left_cell": "A1",
    "active_pane": "topLeft",
    "panes": [
			{
				"sqref": "A1:E1",
				"active_cell": "A1",
				"pane": "topLeft"
			}
		]
	}`)

	// Set Column Style
	xl.SetColStyle(indexSheetName, "A", style.Center(xl))
	xl.SetColStyle(indexSheetName, "B", style.Center(xl))
	xl.SetColStyle(indexSheetName, "C", style.Center(xl))
	xl.SetColStyle(indexSheetName, "D", style.Left(xl))
	xl.SetColStyle(indexSheetName, "E", style.Left(xl))
	xl.SetColWidth(indexSheetName, "D", "E", 45.0)

	// Set Header
	xl.SetCellStr(indexSheetName, "A1", "#")
	xl.SetCellStr(indexSheetName, "B1", "tag")
	xl.SetCellStr(indexSheetName, "C1", "method")
	xl.SetCellStr(indexSheetName, "D1", "path")
	xl.SetCellStr(indexSheetName, "E1", "summary")

	// Set Data
	row := 0
	for path, operations := range swaggerAPI.Paths {
		for operation, detail := range operations {
			row++
			fmt.Printf("OPERATION %s\n", operation)
			fmt.Printf("%v\n", detail)
			xl.SetCellInt(indexSheetName, fmt.Sprintf("%s%d", "A", row+1), row)
			xl.SetCellHyperLink(indexSheetName, fmt.Sprintf("%s%d", "A", row+1), fmt.Sprintf("%d!A1", row), "Location")
			if len(detail.Tags) > 0 {
				xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "B", row+1), detail.Tags[0])
			}
			xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "C", row+1), operation)
			xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "D", row+1), path)
			xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "E", row+1), detail.Summary)
			createAPISheet(xl, path, operation, detail, row)
		}
	}

	err := xl.AddTable(indexSheetName, "A1", fmt.Sprintf("%s%d", "E", row+1), `{
    "table_name": "table",
    "table_style": "TableStyleMedium21",
    "show_first_column": false,
    "show_last_column": false,
    "show_row_stripes": true,
    "show_column_stripes": false
	}`)
	if err != nil {
		log.Fatalln(err)
	}
	return xl
}

func createAPISheet(xl *excelize.File, path, operation string, detail spec.Operation, sheetName int) {
	worksheetName := strconv.Itoa(sheetName)
	xl.NewSheet(worksheetName)
	xl.SetColWidth(worksheetName, "A", "A", 10)
	// xl.SetColWidth(worksheetName, "B", "B", 45.0)

	// xl.MergeCell(worksheetName, "G1", "H2")
	xl.SetCellStr(worksheetName, "G1", "INDEX")
	xl.SetCellHyperLink(worksheetName, "G1", "INDEX!A1", "Location")
	xl.SetCellStyle(worksheetName, "G1", "G1", style.Button(xl))

	row := 1
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Tag")
	if len(detail.Tags) > 0 {
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), detail.Tags[0])
	}
	row++
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Method")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), operation)
	row++
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Path")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), path)
	row++
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Summary")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), detail.Summary)
	row++
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "Description")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), detail.Description)
	row++

	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "REQUEST")
	xl.SetRowHeight(worksheetName, row, 15)
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), style.Title(xl))
	row++

	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "required")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), "parameter")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "type")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), "level")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "data")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), "description")
	row++

	for _, param := range detail.Parameters {
		fmt.Println("param:", param)
		if param.Required {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "O")
		} else {
			xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "X")
		}
		// TODO: Set definitions
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), param.Name)
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), param.In)
		// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), param.Level)
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), param.Type)
		xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), param.Description)
		row++
	}
	row++

	xl.MergeCell(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row))
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "RESPONSE")
	xl.SetRowHeight(worksheetName, row, 15)
	xl.SetCellStyle(worksheetName, fmt.Sprintf("%s%d", "A", row), fmt.Sprintf("%s%d", "F", row), style.Title(xl))
	row++

	// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "A", row), "required")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), "parameter")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), "type")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), "level")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), "data")
	xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), "description")
	row++

	for _, response := range detail.Responses {
		fmt.Println("response:", response)
		fmt.Println("response.Schema:", response.Schema)
		fmt.Println("response.Schema.Ref:", response.Schema.Ref != "")
		fmt.Println("response.Schema.Ref:", response.Schema.Ref)
		// TODO: Set definitions
		// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "B", row), response.Name)
		// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "C", row), response.In)
		// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "D", row), response.Level)
		// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "E", row), response.Type)
		// xl.SetCellStr(worksheetName, fmt.Sprintf("%s%d", "F", row), response.Description)
		row++
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
