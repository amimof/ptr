package ptr

// StrPtr returns pointer to a string.
func Str(val string) *string {
	return &val
}

// BoolPtr returns pointer to an boolean.
func Bool(val bool) *bool {
	return &val
}

// IntPtr returns pointer to an integer.
func Int(val int) *int {
	return &val
}

// Uint8Ptr returns pointer to an uint8.
func Uint8(val uint8) *uint8 {
	return &val
}

// Uint16Ptr returns pointer to an uint16.
func Uint16(val uint16) *uint16 {
	return &val
}

// Float32Ptr returns pointer to a float32.
func Float32(val float32) *float32 {
	return &val
}

// Float64Ptr returns pointer to a float64.
func Float64(val float64) *float64 {
	return &val
}
