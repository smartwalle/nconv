package conv4go

func Int(value interface{}) int {
	if v, ok := value.(int); ok {
		return v
	}
	return int(Float64(value))
}

func Int8(value interface{}) int8 {
	if v, ok := value.(int8); ok {
		return v
	}
	return int8(Float64(value))
}

func Int16(value interface{}) int16 {
	if v, ok := value.(int16); ok {
		return v
	}
	return int16(Float64(value))
}

func Int32(value interface{}) int32 {
	if v, ok := value.(int32); ok {
		return v
	}
	return int32(Float64(value))
}

func Int64(value interface{}) int64 {
	if v, ok := value.(int64); ok {
		return v
	}
	return int64(Float64(value))
}

func Uint(value interface{}) uint {
	if v, ok := value.(uint); ok {
		return v
	}
	return uint(Float64(value))
}

func Uint8(value interface{}) uint8 {
	if v, ok := value.(uint8); ok {
		return v
	}
	return uint8(Float64(value))
}

func Uint16(value interface{}) uint16 {
	if v, ok := value.(uint16); ok {
		return v
	}
	return uint16(Float64(value))
}

func Uint32(value interface{}) uint32 {
	if v, ok := value.(uint32); ok {
		return v
	}
	return uint32(Float64(value))
}

func Uint64(value interface{}) uint64 {
	if v, ok := value.(uint64); ok {
		return v
	}
	return uint64(Float64(value))
}