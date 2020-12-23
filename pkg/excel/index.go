package excel

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/go-openapi/spec"
)

func (xl *Excel) createIndexSheet() error {
	err := xl.File.SetDocProps(&excelize.DocProperties{
		Category:    "OpenAPI",
		Created:     time.Now().Format(time.RFC3339),
		Modified:    time.Now().Format(time.RFC3339),
		Creator:     "Swage",
		Description: "Open API Specification",
		Identifier:  "xlsx",
	})
	if err != nil {
		log.Fatalln(err)
	}
	xl.File.SetSheetName("Sheet1", xl.indexSheetName)
	xl.File.SetPanes(xl.indexSheetName, `{
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
	xl.File.SetColStyle(xl.indexSheetName, "A", xl.Style.Center)
	xl.File.SetColStyle(xl.indexSheetName, "B", xl.Style.Center)
	xl.File.SetColStyle(xl.indexSheetName, "C", xl.Style.Center)
	xl.File.SetColStyle(xl.indexSheetName, "D", xl.Style.Left)
	xl.File.SetColStyle(xl.indexSheetName, "E", xl.Style.Left)
	xl.File.SetColWidth(xl.indexSheetName, "B", "B", 16.4)
	xl.File.SetColWidth(xl.indexSheetName, "D", "E", 60.0)

	// Set Header
	xl.File.SetCellStr(xl.indexSheetName, "A1", "#")
	xl.File.SetCellStr(xl.indexSheetName, "B1", "tag")
	xl.File.SetCellStr(xl.indexSheetName, "C1", "method")
	xl.File.SetCellStr(xl.indexSheetName, "D1", "path")
	xl.File.SetCellStr(xl.indexSheetName, "E1", "summary")

	// Sort a map by path
	paths := make([]string, 0, len(xl.SwaggerSpec.Paths.Paths))
	for path := range xl.SwaggerSpec.Paths.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	// Set Data
	row := 1
	for _, path := range paths {
		operations := xl.SwaggerSpec.Paths.Paths[path]
		if operations.PathItemProps.Get != nil {
			xl.setOperation(row, path, "GET", operations.PathItemProps.Get, xl.SwaggerSpec.Definitions)
			row++
		}
		if operations.PathItemProps.Put != nil {
			xl.setOperation(row, path, "PUT", operations.PathItemProps.Put, xl.SwaggerSpec.Definitions)
			row++
		}
		if operations.PathItemProps.Post != nil {
			xl.setOperation(row, path, "POST", operations.PathItemProps.Post, xl.SwaggerSpec.Definitions)
			row++
		}
		if operations.PathItemProps.Delete != nil {
			xl.setOperation(row, path, "DELETE", operations.PathItemProps.Delete, xl.SwaggerSpec.Definitions)
			row++
		}
		if operations.PathItemProps.Options != nil {
			xl.setOperation(row, path, "OPTIONS", operations.PathItemProps.Options, xl.SwaggerSpec.Definitions)
			row++
		}
		if operations.PathItemProps.Head != nil {
			xl.setOperation(row, path, "HEAD", operations.PathItemProps.Head, xl.SwaggerSpec.Definitions)
			row++
		}
		if operations.PathItemProps.Patch != nil {
			xl.setOperation(row, path, "PATCH", operations.PathItemProps.Patch, xl.SwaggerSpec.Definitions)
			row++
		}
	}

	err = xl.File.AddTable(xl.indexSheetName, "A1", fmt.Sprintf("%s%d", "E", row+1), `{
    "table_name": "table",
    "table_style": "TableStyleMedium21",
    "show_first_column": false,
    "show_last_column": false,
    "show_row_stripes": true,
    "show_column_stripes": false
	}`)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (xl *Excel) setOperation(row int, path, method string, operation *spec.Operation, definitions spec.Definitions) {
	xl.File.SetCellInt(xl.indexSheetName, fmt.Sprintf("%s%d", "A", row+1), row)
	xl.File.SetCellHyperLink(xl.indexSheetName, fmt.Sprintf("%s%d", "A", row+1), fmt.Sprintf("%d!A1", row), "Location")
	xl.File.SetCellStr(xl.indexSheetName, fmt.Sprintf("%s%d", "B", row+1), strings.Join(operation.Tags, ";"))
	xl.File.SetCellStr(xl.indexSheetName, fmt.Sprintf("%s%d", "C", row+1), method)
	xl.File.SetCellStr(xl.indexSheetName, fmt.Sprintf("%s%d", "D", row+1), path)
	xl.File.SetCellStr(xl.indexSheetName, fmt.Sprintf("%s%d", "E", row+1), operation.Summary)
	xl.createAPISheet(path, method, operation, definitions, row)
	row++
}
