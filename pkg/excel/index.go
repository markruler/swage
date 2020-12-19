package excel

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/markruler/swage/pkg/spec"
	"github.com/markruler/swage/pkg/style"
)

var (
	indexSheetName = "INDEX"
)

func createIndexSheet(swaggerAPI *spec.SwaggerAPI) *excelize.File {
	xl := excelize.NewFile()
	err := xl.SetDocProps(&excelize.DocProperties{
		Category: "OpenAPI",
		Created: time.Now().String(),
		Modified: time.Now().String(),
		Creator: "Swage",
		Description: "Open API Specification",
		Identifier: "xlsx",
	})
	if err != nil {
		log.Fatalln(err)
	}
	xl.SetSheetName("Sheet1", indexSheetName)
	xl.SetPanes(indexSheetName, `{
    "freeze": true,
    "split": true,
    "x_split": 1,
    "y_split": 1,
    "top_left_cell": "B2",
    "active_pane": "topLeft",
    "panes": [
			{
				"sqref": "B2",
				"active_cell": "B2",
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

	// Sort a map by path
	paths := make([]string, 0, len(swaggerAPI.Paths))
	for path := range swaggerAPI.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	
	// Set Data
	row := 0
	for _, path := range paths {
		operations := swaggerAPI.Paths[path]
		for operation, detail := range operations {
			row++
			xl.SetCellInt(indexSheetName, fmt.Sprintf("%s%d", "A", row+1), row)
			xl.SetCellHyperLink(indexSheetName, fmt.Sprintf("%s%d", "A", row+1), fmt.Sprintf("%d!A1", row), "Location")
			xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "B", row+1), strings.Join(detail.Tags, ";"))
			xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "C", row+1), operation)
			xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "D", row+1), path)
			xl.SetCellStr(indexSheetName, fmt.Sprintf("%s%d", "E", row+1), detail.Summary)
			createAPISheet(xl, path, operation, detail, swaggerAPI.Definitions, row)
		}
	}

	err = xl.AddTable(indexSheetName, "A1", fmt.Sprintf("%s%d", "E", row+1), `{
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
