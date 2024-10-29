package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var GetField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		GetField = val.Field
	case reflect.Slice:
		numberOfValues = val.Len()
		GetField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(GetField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
