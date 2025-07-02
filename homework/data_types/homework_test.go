package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

// Преобразует 32-битное беззнаковое целое число из формата Big Endian в Little Endian.
// Параметры:
//   - number: Исходное число в формате Big Endian.
//
// Возвращает:
//   - Число в формате Little Endian.
func ConvertToLittleEndian(number uint32) uint32 {
	firstByte := (number & 0xFF000000) >> 24
	secondByte := (number & 0x00FF0000) >> 16
	thirdByte := (number & 0x0000FF00) >> 8
	fourthByte := (number & 0x000000FF)
	number = (fourthByte << 24) | (thirdByte << 16) | (secondByte << 8) | firstByte
	return number
}

func TestConversion(t *testing.T) {
	tests := []struct {
		name   string
		number uint32
		want   uint32
	}{
		{
			name:   "Zero value",
			number: 0x00000000,
			want:   0x00000000,
		},
		{
			name:   "All bits set",
			number: 0xFFFFFFFF,
			want:   0xFFFFFFFF,
		},
		{
			name:   "Interleaved bytes",
			number: 0x00FF00FF,
			want:   0xFF00FF00,
		},
		{
			name:   "Half max value",
			number: 0x0000FFFF,
			want:   0xFFFF0000,
		},
		{
			name:   "Random sequence",
			number: 0x01020304,
			want:   0x04030201,
		},
		{
			name:   "Alternating pattern",
			number: 0x55AA55AA,
			want:   0xAA55AA55,
		},
		{
			name:   "Symmetric number",
			number: 0x12343210,
			want:   0x10323412,
		},
		{
			name:   "Minimal positive",
			number: 0x00000001,
			want:   0x01000000,
		},
		{
			name:   "Near maximum",
			number: 0xFEEDBEEF,
			want:   0xEFBEEDFE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToLittleEndian(tt.number)
			assert.Equalf(t, tt.want, got, "Test failed for input: %d", tt.number)
		})
	}
}
