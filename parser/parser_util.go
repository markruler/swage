package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	oas "github.com/go-openapi/spec"
)

func SortMap(hashmap interface{}) []string {
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

func Enum2string(enums ...interface{}) string {
	var enumSlice []string
	for _, enum := range enums {
		enumSlice = append(enumSlice, fmt.Sprintf("%v", enum))
	}
	enumString := strings.Join(enumSlice, ",")
	return enumString
}

func DefinitionNameFromRef(ref oas.Ref) string {
	url := ref.GetURL()
	if url == nil || url.String() == "" {
		return ""
	}
	lastIndex := strings.LastIndex(url.Fragment, "/")
	return url.Fragment[lastIndex+1:]
}

func checkRequired(required bool) string {
	if required {
		return "O"
	}
	return "X"
}

func extractExample(example interface{}) (string, error) {
	if example == nil {
		return "", errors.New("example is empty")
	}
	b, err := json.MarshalIndent(example, "", "    ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
