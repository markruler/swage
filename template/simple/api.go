package simple

import (
	"errors"
	"strconv"

	"github.com/go-openapi/spec"
)

func (simple *Simple) CreateAPISheet(path, method string, operation *spec.Operation, definitions spec.Definitions, sheetName int) (err error) {
	xl := simple.xl
	if operation == nil {
		return errors.New("operation should not be empty")
	}
	xl.WorkSheetName = strconv.Itoa(sheetName)
	xl.File.NewSheet(xl.WorkSheetName)

	xl.Context.Row = 1
	simple.setAPISheetHeader(path, method, operation)
	simple.setAPISheetRequest(operation)
	if err = simple.setAPISheetResponse(operation); err != nil {
		return err
	}
	return nil
}
