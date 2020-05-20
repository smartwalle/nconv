package conv4go

import (
	"reflect"
	"strconv"
)

func Float32(value interface{}) float32 {
	if v, ok := value.(float32); ok {
		return v
	}
	return float32(Float64(value))
}

func Float64(value interface{}) float64 {
	if v, ok := value.(float64); ok {
		return v
	}
	return floatValue(value)
}

func floatValue(value interface{}) float64 {
	var vValue = reflect.ValueOf(value)
	var vKind = vValue.Kind()

	switch vKind {
	case reflect.Bool:
		var v = vValue.Bool()
		if v {
			return 1.0
		}
		return 0.0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(vValue.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(vValue.Int())
	case reflect.Float32, reflect.Float64:
		return vValue.Float()
	case reflect.String:
		var f, err = strconv.ParseFloat(vValue.String(), 64)
		if err == nil {
			return f
		}
	}
	return 0.0
}
