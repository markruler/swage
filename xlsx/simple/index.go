package simple

import (
	"errors"
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func (simple *Simple) CreateIndexSheet() error {
	xl := simple.xl
	if len(xl.SwageSpec.API) == 0 {
		return errors.New("api is empty")
	}

	err := xl.File.SetDocProps(&excelize.DocProperties{
		Category:    "OpenAPI",
		Created:     time.Now().Format(time.RFC3339),
		Modified:    time.Now().Format(time.RFC3339),
		Creator:     "Swage",
		Description: "Open API Specification",
		Identifier:  "xlsx",
	})
	if err != nil {
		return err
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
	for _, api := range xl.SwageSpec.API {
		xl.File.SetCellInt(xl.IndexSheetName, fmt.Sprintf("%s%d", "A", row+1), row)
		xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "B", row+1), api.Header.Tag)
		xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "C", row+1), api.Header.Method)
		xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "D", row+1), api.Header.Path)
		xl.File.SetCellStr(xl.IndexSheetName, fmt.Sprintf("%s%d", "E", row+1), api.Header.Summary)
		xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "A", row+1), fmt.Sprintf("%d!A1", row), "Location")
		xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "B", row+1), fmt.Sprintf("%d!A1", row), "Location")
		xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "C", row+1), fmt.Sprintf("%d!A1", row), "Location")
		xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "D", row+1), fmt.Sprintf("%d!A1", row), "Location")
		xl.File.SetCellHyperLink(xl.IndexSheetName, fmt.Sprintf("%s%d", "E", row+1), fmt.Sprintf("%d!A1", row), "Location")
		row++
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
		return err
	}
	return nil
}
