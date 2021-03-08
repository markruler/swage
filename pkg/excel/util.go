package excel

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/go-openapi/spec"
)

func sortMap(hashmap interface{}) []string {
	reflection := reflect.ValueOf(hashmap)
	if reflection.Kind() == reflect.Map {
		keys := make([]string, 0, len(reflection.MapKeys()))
		for _, key := range reflection.MapKeys() {
			// interfaceByKey := reflection.MapIndex(key)
			// fmt.Println("reflection:", key.Interface(), interfaceByKey.Interface())
			keys = append(keys, fmt.Sprintf("%v", key.Interface()))
		}
		sort.Strings(keys)
		return keys
	}
	return nil
}

func enum2string(enums ...interface{}) string {
	var enumSlice []string
	for _, enum := range enums {
		enumSlice = append(enumSlice, enum.(string))
	}
	enumString := strings.Join(enumSlice, ",")
	return enumString
}

func (xl *Excel) getParameterFromRef(ref spec.Ref) *spec.Parameter {
	url := ref.GetURL()
	if url == nil || url.String() == "" {
		return nil
	}
	lastIndex := strings.LastIndex(url.Fragment, "/")
	parameterName := url.Fragment[lastIndex+1:]
	if xl.SwaggerSpec == nil || len(xl.SwaggerSpec.Parameters) == 0 {
		return nil
	}
	param := xl.SwaggerSpec.Parameters[parameterName]
	return &param
}

func (xl *Excel) getDefinitionFromRef(ref spec.Ref) (definitionName string, definition *spec.Schema) {
	url := ref.GetURL()
	if url == nil || url.String() == "" {
		return "", nil
	}
	lastIndex := strings.LastIndex(url.Fragment, "/")
	defName := url.Fragment[lastIndex+1:]
	if xl.SwaggerSpec == nil || len(xl.SwaggerSpec.Definitions) == 0 {
		return "", nil
	}
	def := xl.SwaggerSpec.Definitions[defName]
	return defName, &def
}

func (xl *Excel) setCellWithSchema(schemaName, paramType, dataType, description string) {
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), schemaName)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), paramType)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), dataType)
	// xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "E", xl.Context.row), enum)
	// xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "F", xl.Context.row), example)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), description)
}

func (xl *Excel) getParameterSchema(param spec.Parameter) error {
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
		if err := xl.getParameterSchemaRef(param); err != nil {
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

func (xl *Excel) getParameterSchemaRef(param spec.Parameter) error {
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

func (xl *Excel) getResponseSchema(response spec.Response) error {
	if response.Schema.Title != "" {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), response.Schema.Title)
	}

	if response.Schema.Type != nil {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(response.Schema.Type, ","))
	}

	if response.Schema.Type.Contains("array") {
		items := response.Schema.Items
		if items.Schema != nil {
			schema := items.Schema
			xl.setCellWithSchema(schema.Title, "body", strings.Join(response.Schema.Type, ","), response.Description)
			return nil
		}
		for _, schema := range items.Schemas {
			if !reflect.DeepEqual(spec.Ref{}, schema.Ref) {
				definitionName, definition := xl.getDefinitionFromRef(items.Schemas[0].Ref)
				if definition == nil {
					return errors.New("not found response.Schema.Items definition")
				}
				xl.setCellWithSchema(definitionName, "body", "array", response.Description)
				return nil
			}
		}
	}

	if response.Schema.Title != "" {
		xl.setCellWithSchema(response.Schema.Title, "body", "object", response.Description)
		xl.Context.row++
		return nil
	}

	// TODO: refactoring if-hell
	// TODO: write test code
	if response.Schema.Properties != nil {
		for propertyName, propertySchema := range response.Schema.Properties {
			if !reflect.DeepEqual(spec.Ref{}, propertySchema.Ref) {
				definitionName, definition := xl.getDefinitionFromRef(propertySchema.Ref)
				if definition == nil {
					return errors.New("not found definition")
				}
				if propertySchema.Items != nil {
					definitionName, definition = xl.getDefinitionFromRef(propertySchema.Items.Schema.Ref)
					if definition == nil {
						return errors.New("not found definition")
					}
				}
				xl.setCellWithSchema(definitionName, "body", propertyName, propertySchema.Description)
				xl.Context.row++
				return nil
			}
			xl.setCellWithSchema(propertyName, "body", strings.Join(response.Schema.Type, ","), response.Description)
			xl.Context.row++
		}
	}
	return nil
}

func (xl *Excel) getResponseSchemaRef(response spec.Response) error {
	schema, err := spec.ResolveRef(*xl.SwaggerSpec, &response.Schema.Ref)
	if err != nil {
		return err
	}
	if schema == nil {
		return errors.New("not found response.Schema.Ref definition")
	}
	// TODO: handle 'AllOf'
	// if len(schema.AllOf) != 0 {
	// 	for _, oneSchema := range schema.AllOf {
	// 		fmt.Println("oneSchema:", oneSchema)
	// 		schema, err := spec.ResolveResponse(xl.SwaggerSpec, oneSchema.Ref)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println("schema:", schema)
	// 		schemaName, _ := xl.getDefinitionFromRef(oneSchema.Ref)
	// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), schemaName)
	// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
	// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), "object")
	// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), schema.Description)
	// 	}
	// 	return nil
	// }
	schemaName, _ := xl.getDefinitionFromRef(response.Schema.Ref)
	xl.setCellWithSchema(schemaName, "body", "object", response.Description)
	xl.Context.row++
	return nil
}
