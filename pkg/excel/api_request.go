package excel

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-openapi/spec"
)

func (xl *Excel) setAPISheetRequest(operation *spec.Operation) {
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "REQUEST")
	xl.File.MergeCell(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row))
	xl.File.SetRowHeight(xl.Context.worksheetName, xl.Context.row, 20.0)
	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Title)
	xl.Context.row++

	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "required")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), "schema")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "param-type")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "data-type")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), "enum")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "F", xl.Context.row), "example")
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), "description")
	xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "G", xl.Context.row), xl.Style.Column)
	xl.Context.row++

	for _, param := range operation.Parameters {
		xl.File.SetCellStyle(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), fmt.Sprintf("%s%d", "F", xl.Context.row), xl.Style.Center)

		if !reflect.DeepEqual(param.Ref, spec.Ref{}) {
			param = *xl.getParameterFromRef(param.Ref)
		}

		if param.Required {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "O")
		} else {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "X")
		}
		xl.setCellWithSchema(param.Name, param.In, param.Type, param.Description)

		if param.Schema != nil {
			getSchema(xl, param)
		}

		if param.Items != nil && param.Items.Enum != nil {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), enum2string(param.Items.Enum...))
		}

		if param.Enum != nil {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), enum2string(param.Enum...))
		}
		// TODO: remove empty row
		xl.Context.row++
	}
	xl.Context.row++
}

func enum2string(enums ...interface{}) string {
	var enumSlice []string
	for _, enum := range enums {
		enumSlice = append(enumSlice, enum.(string))
	}
	enumString := strings.Join(enumSlice, ",")
	return enumString
}

func getSchema(xl *Excel, param spec.Parameter) error {
	// TODO: write test code
	if param.Schema.Items != nil {
		if param.Schema.Items.Schema != nil {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ","))
		}
		// TODO: Schema's'
		if param.Schema.Items.Schemas != nil {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ","))
		}
	}

	if !reflect.DeepEqual(spec.Ref{}, param.Schema.Ref) {
		if err := getSchemaRef(xl, param); err != nil {
			return err
		}
		// continue
		return nil
	}

	if param.Schema.Properties != nil {
		for k, v := range param.Schema.Properties {
			xl.setCellWithSchema(k, param.In, strings.Join(v.Type, ","), "")
		}
	}

	if param.Schema.Type != nil {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ","))
	}

	if param.Schema.Description != "" {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), param.Schema.Description)
	}

	xl.Context.row++
	return nil
}

func getSchemaRef(xl *Excel, param spec.Parameter) error {
	if strings.Contains(param.Schema.Ref.GetPointer().String(), "definitions") {
		schema, err := spec.ResolveRef(xl.SwaggerSpec, &param.Schema.Ref)
		if err != nil {
			return err
		}

		if param.Required {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "O")
		} else {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "X")
		}

		schemaName, _ := xl.getDefinitionFromRef(param.Schema.Ref)
		xl.setCellWithSchema(schemaName, param.In, strings.Join(schema.Type, ","), param.Description)
		
		xl.Context.row++
	}

	if strings.Contains(param.Schema.Ref.GetPointer().String(), "parameters") {
		schema, err := spec.ResolveParameter(xl.SwaggerSpec, param.Schema.Ref)
		if err != nil {
			return err
		}

		if schema.Required {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "O")
		} else {
			xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "X")
		}

		xl.setCellWithSchema(schema.Name, schema.In, schema.Type, schema.Description)
		xl.Context.row++
	}
	return nil
}
