package excel

import (
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
