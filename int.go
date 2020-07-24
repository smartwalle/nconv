package conv4go

import (
	"reflect"
	"strconv"
	"strings"
)

func Int(value interface{}) int {
	if v, ok := value.(int); ok {
		return v
	}
	return int(intValue(value))
}

func Int8(value interface{}) int8 {
	if v, ok := value.(int8); ok {
		return v
	}
	return int8(intValue(value))
}

func Int16(value interface{}) int16 {
	if v, ok := value.(int16); ok {
		return v
	}
	return int16(intValue(value))
}

func Int32(value interface{}) int32 {
	if v, ok := value.(int32); ok {
		return v
	}
	return int32(intValue(value))
}

func Int64(value interface{}) int64 {
	if v, ok := value.(int64); ok {
		return v
	}
	return intValue(value)
}

func Uint(value interface{}) uint {
	if v, ok := value.(uint); ok {
		return v
	}
	return uint(uintValue(value))
}

func Uint8(value interface{}) uint8 {
	if v, ok := value.(uint8); ok {
		return v
	}
	return uint8(uintValue(value))
}

func Uint16(value interface{}) uint16 {
	if v, ok := value.(uint16); ok {
		return v
	}
	return uint16(uintValue(value))
}

func Uint32(value interface{}) uint32 {
	if v, ok := value.(uint32); ok {
		return v
	}
	return uint32(uintValue(value))
}

func Uint64(value interface{}) uint64 {
	if v, ok := value.(uint64); ok {
		return v
	}
	return uintValue(value)
}

func intValue(value interface{}) int64 {
	var vValue = reflect.ValueOf(value)
	var vKind = vValue.Kind()

	switch vKind {
	case reflect.Bool:
		var v = vValue.Bool()
		if v {
			return 1
		}
		return 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(vValue.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vValue.Int()
	case reflect.Float32, reflect.Float64:
		return int64(vValue.Float())
	case reflect.String:
		var vList = strings.Split(vValue.String(), ".")
		var f, err = strconv.ParseInt(vList[0], 10, 64)
		if err == nil {
			return f
		}
	}
	return 0
}

func uintValue(value interface{}) uint64 {
	var vValue = reflect.ValueOf(value)
	var vKind = vValue.Kind()

	switch vKind {
	case reflect.Bool:
		var v = vValue.Bool()
		if v {
			return 1
		}
		return 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return vValue.Uint()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(vValue.Int())
	case reflect.Float32, reflect.Float64:
		return uint64(vValue.Float())
	case reflect.String:
		var vList = strings.Split(vValue.String(), ".")
		var f, err = strconv.ParseUint(vList[0], 10, 64)
		if err == nil {
			return f
		}
	}
	return 0
}
