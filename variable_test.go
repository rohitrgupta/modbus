package modbus

import (
	"math"
	"testing"
)

func TestVariable(t *testing.T) {

	var tests = []struct {
		name          string
		fmt           string
		endian        Endianness
		offset, scale float32
		want          int64
		reg           []byte
	}{
		{"V01", "uint16", BIG_ENDIAN, 0, 1, 1, []byte{00, 01}},
		{"V02", "uint16", BIG_ENDIAN, 0, 1, 256, []byte{01, 00}},
		{"V03", "uint16", BIG_ENDIAN, 0, 1, 256*256 - 1, []byte{0xFF, 0xFF}},
		{"V04", "uint16", LITTLE_ENDIAN, 0, 1, 1, []byte{01, 00}},
		{"V05", "uint16", LITTLE_ENDIAN, 0, 1, 256, []byte{00, 01}},
		{"V06", "int16", BIG_ENDIAN, 0, 1, 1, []byte{00, 01}},
		{"V07", "int16", BIG_ENDIAN, 0, 1, 256, []byte{01, 00}},
		{"V08", "int16", BIG_ENDIAN, 0, 1, -1, []byte{0xFF, 0xFF}},
		{"V11", "uint32", BIG_ENDIAN, 0, 1, 1, []byte{00, 00, 00, 01}},
		{"V12", "uint32", BIG_ENDIAN, 0, 1, 256 * 1, []byte{00, 00, 01, 00}},
		{"V13", "uint32", BIG_ENDIAN, 0, 1, 256 * 256 * 1, []byte{00, 01, 00, 00}},
		{"V14", "uint32", BIG_ENDIAN, 0, 1, 256 * 256 * 256 * 1, []byte{01, 00, 00, 00}},
		{"V15", "uint32", BIG_ENDIAN, 0, 1, 256*256*256*256 - 1, []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{"V16", "uint32", LITTLE_ENDIAN, 0, 1, 1, []byte{01, 00, 00, 00}},
		{"V17", "uint32", LITTLE_ENDIAN, 0, 1, 256 * 1, []byte{00, 01, 00, 00}},
		{"V18", "uint32", LITTLE_ENDIAN, 0, 1, 256 * 256 * 1, []byte{00, 00, 01, 00}},
		{"V19", "uint32", LITTLE_ENDIAN, 0, 1, 256 * 256 * 256 * 1, []byte{00, 00, 00, 01}},
		{"V20", "int32", BIG_ENDIAN, 0, 1, -1, []byte{0xFF, 0xFF, 0xFF, 0xFF}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var1 := MbVar{name: tt.name, fmt: tt.fmt, offset: tt.offset, scale: tt.scale, endian: tt.endian}
			var1.SetReg(tt.reg)

			if var1.valueType != VALUE_TYPE_INT {
				t.Errorf("got value type %d want %d", var1.valueType, VALUE_TYPE_INT)
			}
			if var1.valueInt != tt.want {
				t.Errorf("got value %d want %d", var1.valueInt, tt.want)

			}
		})
	}
}

const float64EqualityThreshold = 1e-3

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
func TestVariableFloat(t *testing.T) {

	var tests = []struct {
		name          string
		fmt           string
		endian        Endianness
		offset, scale float32
		want          float64
		reg           []byte
	}{
		{"V01", "uint16", BIG_ENDIAN, 1, 1, 2, []byte{00, 01}},
		{"V02", "uint16", BIG_ENDIAN, 10, 1, 11, []byte{00, 01}},
		{"V03", "uint16", BIG_ENDIAN, 0, 10, 10, []byte{00, 01}},
		{"V04", "uint16", BIG_ENDIAN, 0, 0.1, 0.2, []byte{00, 02}},
		{"V05", "uint16", BIG_ENDIAN, 1, 0.1, 1.3, []byte{00, 03}},
		{"V06", "uint16", BIG_ENDIAN, 100, 0.1, 100.4, []byte{00, 04}},
		{"V07", "float32", BIG_ENDIAN, 0, 1, 0, []byte{0, 0, 0, 0}},
		{"V08", "float32", BIG_ENDIAN, 0, 1, 1, []byte{0x3f, 0x80, 0x00, 0x00}},
		{"V09", "float32", BIG_ENDIAN, 0, 1, 2, []byte{0x40, 0, 0, 0}},
		{"V09", "float32", BIG_ENDIAN, 0, 1, 0.5, []byte{0x3f, 0, 0, 0}},
		{"V09", "float32", BIG_ENDIAN, 0, 1, -1, []byte{0xbf, 0x80, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var1 := MbVar{name: tt.name, fmt: tt.fmt, offset: tt.offset, scale: tt.scale, endian: tt.endian}
			var1.SetReg(tt.reg)

			if var1.valueType != VALUE_TYPE_FLOAT {
				t.Errorf("got value type %d want %d", var1.valueType, VALUE_TYPE_FLOAT)
			}
			if !almostEqual(var1.valueFloat, tt.want) {
				t.Errorf("got value %f want %f", var1.valueFloat, tt.want)

			}
		})
	}
}
