package simple

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/markruler/swage/parser"
)

func (simple *Simple) parameterFromRef(ref spec.Ref) *spec.Parameter {
	xl := simple.xl
	if xl.SwaggerSpec == nil || len(xl.SwaggerSpec.Parameters) == 0 {
		return nil
	}
	name := parser.DefinitionNameFromRef(ref)
	param := xl.SwaggerSpec.Parameters[name]
	return &param
}

func (simple *Simple) definitionFromRef(ref spec.Ref) (definitionName string, definition *spec.Schema) {
	xl := simple.xl
	if xl.SwaggerSpec == nil || len(xl.SwaggerSpec.Definitions) == 0 {
		return "", nil
	}
	name := parser.DefinitionNameFromRef(ref)
	def := xl.SwaggerSpec.Definitions[name]
	return name, &def
}

func (simple *Simple) setCellWithSchema(schemaName, paramType, dataType, example, description string) {
	xl := simple.xl
	if example == "null" {
		example = ""
	}

	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "B", xl.Context.Row), schemaName)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "C", xl.Context.Row), paramType)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), dataType)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), example)
	xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), description)
}

func (simple *Simple) parameterSchema(param spec.Parameter) error {
	xl := simple.xl
	// FIXME: converting items
	// if param.Schema.Items != nil {
	// 	if param.Schema.Items.Schema != nil {
	// 		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), strings.Join(param.Schema.Type, ","))
	// 	}
	// 	if param.Schema.Items.Schemas != nil {
	// 		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), strings.Join(param.Schema.Type, ","))
	// 	}
	// }

	if !reflect.DeepEqual(spec.Ref{}, param.Schema.Ref) {
		if err := simple.parameterSchemaRef(param); err != nil {
			return err
		}
		// continue
		return nil
	}

	if param.Schema.Properties != nil {
		for k, v := range param.Schema.Properties {
			simple.setCellWithSchema(k, param.In, strings.Join(v.Type, ","), "", "")
		}
		return nil
	}

	if param.Schema.Type != nil {
		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), strings.Join(param.Schema.Type, ","))
	}

	if param.Schema.Example != nil {
		b, err := json.MarshalIndent(param.Schema.Example, "", "    ")
		if err != nil {
			return err
		}
		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "F", xl.Context.Row), string(b))
	}

	if param.Schema.Description != "" {
		xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "G", xl.Context.Row), param.Schema.Description)
	}

	return nil
}

func (simple *Simple) parameterSchemaRef(param spec.Parameter) error {
	if strings.Contains(param.Schema.Ref.GetPointer().String(), "definitions") {
		schema, err := spec.ResolveRef(simple.xl.SwaggerSpec, &param.Schema.Ref)
		if err != nil {
			return err
		}
		simple.checkRequired(param.Required)

		schemaName, _ := simple.definitionFromRef(param.Schema.Ref)

		b, err := json.MarshalIndent(schema.Example, "", "    ")
		if err != nil {
			return err
		}
		simple.setCellWithSchema(schemaName, param.In, strings.Join(schema.Type, ","), string(b), param.Description)
		return nil
	}

	if strings.Contains(param.Schema.Ref.GetPointer().String(), "parameters") {
		schema, err := spec.ResolveParameter(simple.xl.SwaggerSpec, param.Schema.Ref)
		if err != nil {
			return err
		}
		simple.checkRequired(schema.Required)

		b, err := json.MarshalIndent(schema.Example, "", "    ")
		if err != nil {
			return err
		}
		simple.setCellWithSchema(schema.Name, schema.In, schema.Type, string(b), schema.Description)
	}
	return nil
}

func (simple *Simple) responseSchema(response spec.Response) error {
	xl := simple.xl

	if response.Schema.Type != nil {
		simple.xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "C", xl.Context.Row), "body")
		simple.xl.File.SetCellStr(xl.WorkSheetName, fmt.Sprintf("%s%d", "D", xl.Context.Row), strings.Join(response.Schema.Type, ","))
	}

	if response.Schema.Type.Contains("array") {
		if err := simple.arrayDefinitionFromSchemaRef(response); err != nil {
			return err
		}
	}

	if response.Schema.Title != "" {
		b, err := json.MarshalIndent(response.Schema.Example, "", "    ")
		if err != nil {
			return err
		}
		simple.setCellWithSchema(response.Schema.Title, "body", "object", string(b), response.Description)
		xl.Context.Row++
		return nil
	}

	if response.Schema.Properties != nil {
		if err := simple.propDefinitionFromSchemaRef(response); err != nil {
			return err
		}
	}
	return nil
}

func (simple *Simple) arrayDefinitionFromSchemaRef(response spec.Response) error {
	items := response.Schema.Items
	if items.Schema != nil {
		schema := items.Schema
		simple.setCellWithSchema(schema.Title, "body", strings.Join(response.Schema.Type, ","), "", response.Description)
		return nil
	}
	for _, schema := range items.Schemas {
		if !reflect.DeepEqual(spec.Ref{}, schema.Ref) {
			definitionName, definition := simple.definitionFromRef(items.Schemas[0].Ref)
			if definition == nil {
				return errors.New("not found definition")
			}
			b, err := json.MarshalIndent(response.Schema.Example, "", "    ")
			if err != nil {
				return err
			}
			simple.setCellWithSchema(definitionName, "body", "array", string(b), response.Description)
			return nil
		}
	}
	return nil
}

func (simple *Simple) propDefinitionFromSchemaRef(response spec.Response) error {
	if reflect.DeepEqual(spec.Response{}, response) {
		return errors.New("response is empty")
	}

	xl := simple.xl
	for propertyName, propertySchema := range response.Schema.Properties {
		if !reflect.DeepEqual(spec.Ref{}, propertySchema.Ref) {
			definitionName, definition := simple.definitionFromRef(propertySchema.Ref)
			if definition == nil {
				return errors.New("not found definition")
			}
			if propertySchema.Items != nil {
				definitionName, definition = simple.definitionFromRef(propertySchema.Items.Schema.Ref)
				if definition == nil {
					return errors.New("not found definition")
				}
			}
			simple.setCellWithSchema(definitionName, "body", propertyName, "", propertySchema.Description)
			xl.Context.Row++
			return nil
		}
		simple.setCellWithSchema(propertyName, "body", strings.Join(response.Schema.Type, ","), "", response.Description)
		xl.Context.Row++
	}
	return nil
}

func (simple *Simple) responseSchemaRef(response spec.Response) error {
	schema, err := spec.ResolveRef(*simple.xl.SwaggerSpec, &response.Schema.Ref)
	if err != nil {
		return err
	}
	if schema == nil {
		return errors.New("not found response.Schema.Ref definition")
	}

	schemaName, _ := simple.definitionFromRef(response.Schema.Ref)

	b, err := json.MarshalIndent(schema.Example, "", "    ")
	if err != nil {
		return err
	}
	simple.setCellWithSchema(schemaName, "body", "object", string(b), response.Description)
	simple.xl.Context.Row++
	return nil
}

func (simple *Simple) checkRequired(required bool) {
	if required {
		simple.xl.File.SetCellStr(simple.xl.WorkSheetName, fmt.Sprintf("%s%d", "A", simple.xl.Context.Row), "O")
	} else {
		simple.xl.File.SetCellStr(simple.xl.WorkSheetName, fmt.Sprintf("%s%d", "A", simple.xl.Context.Row), "X")
	}
}
