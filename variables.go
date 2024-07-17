package modbus

import (
	"encoding/binary"
	"math"
)

type Variable interface {
	SetReg(regBs []byte)
	GetValue() float32
}

type ValueType uint

const (
	VALUE_TYPE_INT   ValueType = 1
	VALUE_TYPE_FLOAT ValueType = 2
)

type MbVar struct {
	Name       string
	Fmt        string
	Endian     Endianness
	Reg        []byte
	ValueFloat float64
	ValueInt   int64
	Offset     float32
	Scale      float32
	ValueType  ValueType
}

func (r *MbVar) SetReg(reg []byte) {
	r.Reg = reg
	switch r.Fmt {
	case "uint16":
		r.decodeUint16()
	case "int16":
		r.decodeInt16()
	case "uint32":
		r.decodeUint32()
	case "int32":
		r.decodeInt32()
	case "float32":
		r.decodeFloat32()
	}
}

// Decodes uint16 value from register bytes.
// if scale or offset are set, the value is converted to float32.
// otherwise, the value is converted to int64.
func (r *MbVar) decodeUint16() {
	var value uint16
	switch r.Endian {
	case BIG_ENDIAN:
		value = binary.BigEndian.Uint16(r.Reg)
	case LITTLE_ENDIAN:
		value = binary.LittleEndian.Uint16(r.Reg)
	}
	if r.Scale != 1 || r.Offset != 0 {
		r.ValueType = VALUE_TYPE_FLOAT
		r.ValueFloat = float64(float32(value)*r.Scale + r.Offset)
	} else {
		r.ValueType = VALUE_TYPE_INT
		r.ValueInt = int64(value)
	}
}

// Decodes int16 value from register bytes.
// if scale or offset are set, the value is converted to float32.
// otherwise, the value is converted to int64.
func (r *MbVar) decodeInt16() {
	var value int16
	switch r.Endian {
	case BIG_ENDIAN:
		value = int16(binary.BigEndian.Uint16(r.Reg))
	case LITTLE_ENDIAN:
		value = int16(binary.LittleEndian.Uint16(r.Reg))
	}
	if r.Scale != 1 || r.Offset != 0 {
		r.ValueType = VALUE_TYPE_FLOAT
		r.ValueFloat = float64(float32(value)*r.Scale + r.Offset)
	} else {
		r.ValueType = VALUE_TYPE_INT
		r.ValueInt = int64(value)
	}
}

// Decodes uint32 value from register bytes.
// if scale or offset are set, the value is converted to float32.
// otherwise, the value is converted to int64.
func (r *MbVar) decodeUint32() {
	var value uint32
	switch r.Endian {
	case BIG_ENDIAN:
		value = binary.BigEndian.Uint32(r.Reg)
	case LITTLE_ENDIAN:
		value = binary.LittleEndian.Uint32(r.Reg)
	}
	if r.Scale != 1 || r.Offset != 0 {
		r.ValueType = VALUE_TYPE_FLOAT
		r.ValueFloat = float64(float32(value)*r.Scale + r.Offset)
	} else {
		r.ValueType = VALUE_TYPE_INT
		r.ValueInt = int64(value)
	}
}

// Decodes int32 value from register bytes.
// if scale or offset are set, the value is converted to float32.
// otherwise, the value is converted to int64.
func (r *MbVar) decodeInt32() {
	var value int32
	switch r.Endian {
	case BIG_ENDIAN:
		value = int32(binary.BigEndian.Uint32(r.Reg))
	case LITTLE_ENDIAN:
		value = int32(binary.LittleEndian.Uint32(r.Reg))
	}
	if r.Scale != 1 || r.Offset != 0 {
		r.ValueType = VALUE_TYPE_FLOAT
		r.ValueFloat = float64(float32(value)*r.Scale + r.Offset)
	} else {
		r.ValueType = VALUE_TYPE_INT
		r.ValueInt = int64(value)
	}
}

// Decodes float32 value from register bytes.
// scale or offset are ignored, and the value is converted to float32.
func (r *MbVar) decodeFloat32() {
	r.ValueType = VALUE_TYPE_FLOAT
	switch r.Endian {
	case BIG_ENDIAN:
		r.ValueFloat = float64(math.Float32frombits(binary.BigEndian.Uint32(r.Reg)))
	case LITTLE_ENDIAN:
		r.ValueFloat = float64(math.Float32frombits(binary.LittleEndian.Uint32(r.Reg)))
	}
}
