package utils

import (
	"reflect"
	"strings"
)

// StructToMap convert nested struct to map[string]interface{}
func StructToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}

	// Get field type
	v := reflect.TypeOf(item)

	// Get field value
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	// Point to actual value if is a pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		// Get field name from json tag
		tagStr := v.Field(i).Tag.Get("json")
		jsonTags := strings.Split(tagStr, ",")
		tag := jsonTags[0]

		// Parse value to interface{}
		field := reflectValue.Field(i).Interface()

		// Check if ignore
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				// If nested, do recursion
				res[tag] = StructToMap(field)
			} else {
				// Not nested, set value
				res[tag] = field
			}
		}
	}
	return res
}
