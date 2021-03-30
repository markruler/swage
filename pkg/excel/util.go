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

func (xl *Excel) parameterFromRef(ref spec.Ref) *spec.Parameter {
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

func (xl *Excel) definitionFromRef(ref spec.Ref) (definitionName string, definition *spec.Schema) {
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
	// FIXME:
	// xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "F", xl.Context.row), example)
	xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), description)
}

func (xl *Excel) parameterSchema(param spec.Parameter) error {
	// FIXME:
	// if param.Schema.Items != nil {
	// 	if param.Schema.Items.Schema != nil {
	// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ","))
	// 	}
	// 	if param.Schema.Items.Schemas != nil {
	// 		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ","))
	// 	}
	// }

	if !reflect.DeepEqual(spec.Ref{}, param.Schema.Ref) {
		if err := xl.parameterSchemaRef(param); err != nil {
			return err
		}
		// continue
		return nil
	}

	if param.Schema.Properties != nil {
		for k, v := range param.Schema.Properties {
			xl.setCellWithSchema(k, param.In, strings.Join(v.Type, ","), "")
		}
		return nil
	}

	if param.Schema.Type != nil {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(param.Schema.Type, ","))
	}

	if param.Schema.Description != "" {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "G", xl.Context.row), param.Schema.Description)
	}

	return nil
}

func (xl *Excel) parameterSchemaRef(param spec.Parameter) error {
	if strings.Contains(param.Schema.Ref.GetPointer().String(), "definitions") {
		schema, err := spec.ResolveRef(xl.SwaggerSpec, &param.Schema.Ref)
		if err != nil {
			return err
		}
		xl.checkRequired(param.Required)

		schemaName, _ := xl.definitionFromRef(param.Schema.Ref)
		xl.setCellWithSchema(schemaName, param.In, strings.Join(schema.Type, ","), param.Description)
		return nil
	}

	if strings.Contains(param.Schema.Ref.GetPointer().String(), "parameters") {
		schema, err := spec.ResolveParameter(xl.SwaggerSpec, param.Schema.Ref)
		if err != nil {
			return err
		}
		xl.checkRequired(schema.Required)

		xl.setCellWithSchema(schema.Name, schema.In, schema.Type, schema.Description)
	}
	return nil
}

func (xl *Excel) responseSchema(response spec.Response) error {
	if response.Schema.Title != "" {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "B", xl.Context.row), response.Schema.Title)
	}

	if response.Schema.Type != nil {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "C", xl.Context.row), "body")
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "D", xl.Context.row), strings.Join(response.Schema.Type, ","))
	}

	if response.Schema.Type.Contains("array") {
		if err := xl.arrayDefinitionFromSchemaRef(response); err != nil {
			return err
		}
	}

	if response.Schema.Title != "" {
		xl.setCellWithSchema(response.Schema.Title, "body", "object", response.Description)
		xl.Context.row++
		return nil
	}

	if response.Schema.Properties != nil {
		if err := xl.propDefinitionFromSchemaRef(response); err != nil {
			return err
		}
	}
	return nil
}

func (xl *Excel) arrayDefinitionFromSchemaRef(response spec.Response) error {
	items := response.Schema.Items
	if items.Schema != nil {
		schema := items.Schema
		xl.setCellWithSchema(schema.Title, "body", strings.Join(response.Schema.Type, ","), response.Description)
		return nil
	}
	for _, schema := range items.Schemas {
		if !reflect.DeepEqual(spec.Ref{}, schema.Ref) {
			definitionName, definition := xl.definitionFromRef(items.Schemas[0].Ref)
			if definition == nil {
				return errors.New("not found definition")
			}
			xl.setCellWithSchema(definitionName, "body", "array", response.Description)
			return nil
		}
	}
	return nil
}

func (xl *Excel) propDefinitionFromSchemaRef(response spec.Response) error {
	if reflect.DeepEqual(spec.Response{}, response) {
		return errors.New("response is empty")
	}
	for propertyName, propertySchema := range response.Schema.Properties {
		if !reflect.DeepEqual(spec.Ref{}, propertySchema.Ref) {
			definitionName, definition := xl.definitionFromRef(propertySchema.Ref)
			if definition == nil {
				return errors.New("not found definition")
			}
			if propertySchema.Items != nil {
				definitionName, definition = xl.definitionFromRef(propertySchema.Items.Schema.Ref)
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
	return nil
}

func (xl *Excel) responseSchemaRef(response spec.Response) error {
	schema, err := spec.ResolveRef(*xl.SwaggerSpec, &response.Schema.Ref)
	if err != nil {
		return err
	}
	if schema == nil {
		return errors.New("not found response.Schema.Ref definition")
	}

	schemaName, _ := xl.definitionFromRef(response.Schema.Ref)
	xl.setCellWithSchema(schemaName, "body", "object", response.Description)
	xl.Context.row++
	return nil
}

func (xl *Excel) checkRequired(required bool) {
	if required {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "O")
	} else {
		xl.File.SetCellStr(xl.Context.worksheetName, fmt.Sprintf("%s%d", "A", xl.Context.row), "X")
	}
}
