package conv4go

import (
	"fmt"
	"reflect"
	"strconv"
)

func String(value interface{}) string {
	if v, ok := value.(string); ok {
		return v
	}
	return stringValue(value)
}

func stringValue(value interface{}) string {
	var vValue = reflect.ValueOf(value)
	var vKind = vValue.Kind()

	switch vKind {
	case reflect.Bool:
		return strconv.FormatBool(vValue.Bool())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(vValue.Uint(), 10)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(vValue.Int(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(vValue.Float(), 'f', 6, 32)
	case reflect.String:
		return vValue.String()
	}
	return fmt.Sprintf("%v", value)
}
