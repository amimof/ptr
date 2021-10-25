package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPointers tests all the above pointer methods
func TestPointers(t *testing.T) {
	s := "test"
	b := true
	i := int(123)
	ui8 := uint8(123)
	ui16 := uint16(123)
	f32 := float32(1.234)
	f64 := float64(1.234)

	assert.True(t, *Str(s) == s)
	assert.True(t, *Bool(b) == b)
	assert.True(t, *Int(i) == i)
	assert.True(t, *Uint8(ui8) == ui8)
	assert.True(t, *Uint16(ui16) == ui16)
	assert.True(t, *Float32(f32) == f32)
	assert.True(t, *Float64(f64) == f64)
}
