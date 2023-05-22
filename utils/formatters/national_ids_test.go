package formatters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatRutWithoutDots(t *testing.T) {
	rut := "12.345.678-9"
	result := FormatRutWithoutDots(rut)

	assert.Equalf(t, "12345678-9", result, "Expected %s, but got %s for input %s", "12345678-9", result, rut)
}

func TestRemoveDotsAndHyphen(t *testing.T) {
	rut := "12.345.678-9"
	result := RemoveDotsAndHyphen(rut)

	assert.Equalf(t, "12345678", result, "Expected %s, but got %s for input %s", "12345678", result, rut)
}

func TestRutClean(t *testing.T) {
	rut := "12.345.678-9"
	result := rutClean(rut)

	assert.Equalf(t, "123456789", result, "Expected %s, but got %s for input %s", "123456789", result, rut)
}
