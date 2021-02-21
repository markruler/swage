package excel

import (
	"errors"
	"strconv"

	"github.com/go-openapi/spec"
)

func (xl *Excel) createAPISheet(path, method string, operation *spec.Operation, definitions spec.Definitions, sheetName int) (err error) {
	if operation == nil {
		return errors.New("Operation should not be empty")
	}
	xl.Context.worksheetName = strconv.Itoa(sheetName)
	xl.File.NewSheet(xl.Context.worksheetName)

	xl.Context.row = 1
	xl.setAPISheetHeader(path, method, operation)
	xl.setAPISheetRequest(operation)
	if err = xl.setAPISheetResponse(operation); err != nil {
		return err
	}
	return nil
}
