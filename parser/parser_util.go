package parser

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
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
		enumSlice = append(enumSlice, enum.(string))
	}
	enumString := strings.Join(enumSlice, ",")
	return enumString
}
