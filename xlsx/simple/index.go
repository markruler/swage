package simple

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/go-openapi/spec"
	"github.com/markruler/swage/parser"
)

func (simple *Simple) CreateIndexSheet() error {
	xl := simple.xl
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
	xl.File.SetSheetName("Sheet1", xl.IndexSheetName)
	xl.File.SetPanes(xl.IndexSheetName, `{
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
	xl.File.SetColStyle(xl.IndexSheetName, "A", xl.Style.Center)
	xl.File.SetColStyle(xl.IndexSheetName, "B", xl.Style.Center)
	xl.File.SetColStyle(xl.IndexSheetName, "C", xl.Style.Center)
	xl.File.SetColStyle(xl.IndexSheetName, "D", xl.Style.Left)
	xl.File.SetColStyle(xl.IndexSheetName, "E", xl.Style.Left)
	xl.File.SetColWidth(xl.IndexSheetName, "B", "B", 23.9)
	xl.File.SetColWidth(xl.IndexSheetName, "D", "E", 60.0)

	// Set Header
	xl.File.SetCellStr(xl.IndexSheetName, "A1", "#")
	xl.File.SetCellStr(xl.IndexSheetName, "B1", "tag")
	xl.File.SetCellStr(xl.IndexSheetName, "C1", "method")
	xl.File.SetCellStr(xl.IndexSheetName, "D1", "path")
	xl.File.SetCellStr(xl.IndexSheetName, "E1", "summary")

	row := 1
	// FIXME: refactor
	// log.Println(len(xl.SwageSpec.API))
	// for _, api := range xl.SwageSpec.API {
	// 	// log.Println(api)
	// 	xl.File.SetCellInt(xl.IndexSheetName, fmt.Sprintf("%s%d", "A", row+1), row)
	// 	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "B", row+1), api.Header.Tag)
	// 	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "C", row+1), api.Header.Method)
	// 	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "D", row+1), api.Header.Path)
	// 	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "E", row+1), api.Header.Summary)
	// 	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "A", row+1), fmt.Sprintf("%d!A1", row), "Location")
	// 	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "B", row+1), fmt.Sprintf("%d!A1", row), "Location")
	// 	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "C", row+1), fmt.Sprintf("%d!A1", row), "Location")
	// 	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "D", row+1), fmt.Sprintf("%d!A1", row), "Location")
	// 	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "E", row+1), fmt.Sprintf("%d!A1", row), "Location")
	// 	row++
	// }

	// TODO: remove
	paths := parser.SortMap(xl.SwaggerSpec.Paths.Paths)
	for _, path := range paths {
		operations := xl.SwaggerSpec.Paths.Paths[path]
		if operations.PathItemProps.Get != nil {
			row, err = simple.setOperation(row, path, "GET", operations.PathItemProps.Get, xl.SwaggerSpec.Definitions)
			if err != nil {
				return err
			}
		}
		if operations.PathItemProps.Put != nil {
			row, err = simple.setOperation(row, path, "PUT", operations.PathItemProps.Put, xl.SwaggerSpec.Definitions)
			if err != nil {
				return err
			}
		}
		if operations.PathItemProps.Post != nil {
			row, err = simple.setOperation(row, path, "POST", operations.PathItemProps.Post, xl.SwaggerSpec.Definitions)
			if err != nil {
				return err
			}
		}
		if operations.PathItemProps.Delete != nil {
			row, err = simple.setOperation(row, path, "DELETE", operations.PathItemProps.Delete, xl.SwaggerSpec.Definitions)
			if err != nil {
				return err
			}
		}
		if operations.PathItemProps.Options != nil {
			row, err = simple.setOperation(row, path, "OPTIONS", operations.PathItemProps.Options, xl.SwaggerSpec.Definitions)
			if err != nil {
				return err
			}
		}
		if operations.PathItemProps.Head != nil {
			row, err = simple.setOperation(row, path, "HEAD", operations.PathItemProps.Head, xl.SwaggerSpec.Definitions)
			if err != nil {
				return err
			}
		}
		if operations.PathItemProps.Patch != nil {
			row, err = simple.setOperation(row, path, "PATCH", operations.PathItemProps.Patch, xl.SwaggerSpec.Definitions)
			if err != nil {
				return err
			}
		}
	}

	err = xl.File.AddTable(xl.IndexSheetName, "A1", fmt.Sprintf("%s%d", "E", row), `{
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

func (simple *Simple) setOperation(row int, path, method string, operation *spec.Operation, definitions spec.Definitions) (int, error) {
	xl := simple.xl
	xl.File.SetCellInt(xl.IndexSheetName, fmt.Sprintf("%s%d", "A", row+1), row)
	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "B", row+1), strings.Join(operation.Tags, ";"))
	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "C", row+1), method)
	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "D", row+1), path)
	xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "E", row+1), operation.Summary)
	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "A", row+1), fmt.Sprintf("%d!A1", row), "Location")
	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "B", row+1), fmt.Sprintf("%d!A1", row), "Location")
	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "C", row+1), fmt.Sprintf("%d!A1", row), "Location")
	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "D", row+1), fmt.Sprintf("%d!A1", row), "Location")
	xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "E", row+1), fmt.Sprintf("%d!A1", row), "Location")

	if err := simple.CreateAPISheet(path, method, operation, definitions, row); err != nil {
		return 0, err
	}
	return row + 1, nil
}
