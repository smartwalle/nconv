package conv4go

import "reflect"

func Bool(value interface{}) bool {
	if v, ok := value.(bool); ok {
		return v
	}

	var vValue = reflect.ValueOf(value)
	var vKind = vValue.Kind()

	switch vKind {
	case reflect.String:
		var v = vValue.String()
		if v == "true" || v == "yes" || v == "on" || v == "t" || v == "y" || v == "1" {
			return true
		}
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if vValue.Int() == 1 {
			return true
		}
		return false
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if vValue.Uint() == 1 {
			return true
		}
		return false
	case reflect.Float32, reflect.Float64:
		if vValue.Float() > 0.9990 {
			return true
		}
		return false
	case reflect.Bool:
		return vValue.Bool()
	}
	return false
}
