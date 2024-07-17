package modbus

import (
	"testing"
)

func TestPayloadSize(t *testing.T) {
	var tests = []struct {
		name    string
		fmt     string
		address int
		start   int
		size    int
	}{
		{"T01", "uint16", 100, 100, 2},
		{"T02", "int16", 10, 10, 2},
		{"T03", "uint32", 10, 10, 4},
		{"T04", "int32", 20, 20, 4},
		{"T05", "float32", 20, 20, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := NewMbPayload()
			mp.AddVariable(tt.start, MbVar{name: tt.name, fmt: tt.fmt, offset: 0, scale: 1, endian: BIG_ENDIAN})
			if mp.start != tt.start {
				t.Errorf("got start %d want %d", mp.start, tt.start)
			}
			if mp.size != tt.size {
				t.Errorf("got size %d want %d", mp.size, tt.size)
			}
		})
	}
}

func TestPayloadSizeM1(t *testing.T) {
	mp := NewMbPayload()
	start1 := 100
	size1 := 2
	mp.AddVariable(start1, MbVar{name: "T01", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	if mp.start != start1 {
		t.Errorf("got start %d want %d", mp.start, start1)
	}
	if mp.size != size1 {
		t.Errorf("got size %d want %d", mp.size, size1)
	}
	size2 := 22
	mp.AddVariable(110, MbVar{name: "T02", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	if mp.start != start1 {
		t.Errorf("got start %d want %d", mp.start, start1)
	}
	if mp.size != size2 {
		t.Errorf("got size %d want %d", mp.size, size2)
	}
	mp.AddVariable(104, MbVar{name: "T03", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	if mp.size != size2 {
		t.Errorf("got size %d want %d", mp.size, size2)
	}
	start3 := 90
	size3 := 42
	mp.AddVariable(90, MbVar{name: "T04", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	if mp.start != start3 {
		t.Errorf("got start %d want %d", mp.start, start3)
	}
	if mp.size != size3 {
		t.Errorf("got size %d want %d", mp.size, size3)
	}
}

func TestPayloadSizeM2(t *testing.T) {
	mp := NewMbPayload()
	start1 := 90
	size1 := 42
	mp.AddVariable(90, MbVar{name: "T01", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(100, MbVar{name: "T02", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(104, MbVar{name: "T03", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(110, MbVar{name: "T04", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	if mp.start != start1 {
		t.Errorf("got start %d want %d", mp.start, start1)
	}
	if mp.size != size1 {
		t.Errorf("got size %d want %d", mp.size, size1)
	}

}

func TestPayloadRegToVar1(t *testing.T) {

	mp := NewMbPayload()
	mp.AddVariable(0, MbVar{name: "T01", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(1, MbVar{name: "T02", fmt: "uint16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(2, MbVar{name: "T03", fmt: "int32", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(4, MbVar{name: "T04", fmt: "uint32", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(6, MbVar{name: "T05", fmt: "float32", offset: 0, scale: 1, endian: BIG_ENDIAN})
	vars := mp.regToVar(0, []byte{0, 1, 0, 2, 0, 0, 0, 3, 0, 0, 0, 4, 0x3f, 0, 0, 0})
	if mp.vars[0].valueInt != 1 {
		t.Errorf("got value %d want %d", mp.vars[0].valueInt, 1)
	}
	if mp.vars[1].valueInt != 2 {
		t.Errorf("got value %d want %d", mp.vars[1].valueInt, 2)
	}
	if mp.vars[2].valueInt != 3 {
		t.Errorf("got value %d want %d", mp.vars[2].valueInt, 3)
	}
	if mp.vars[4].valueInt != 4 {
		t.Errorf("got value %d want %d", mp.vars[3].valueInt, 4)
	}
	if mp.vars[6].valueFloat != 0.5 {
		t.Errorf("got value %f want %f", mp.vars[3].valueFloat, 0.5)
	}
	if vars[0].Value != int64(1) {
		t.Errorf("got value %d want %d", vars[0].Value, 1)
	}
	if vars[1].Value != int64(2) {
		t.Errorf("got value %d want %d", vars[1].Value, 2)
	}
	if vars[2].Value != int64(3) {
		t.Errorf("got value %d want %d", vars[2].Value, 3)
	}
	if vars[3].Value != int64(4) {
		t.Errorf("got value %d want %d", vars[3].Value, 4)
	}
	if vars[4].Value != 0.5 {
		t.Errorf("got value %f want %f", vars[4].Value, 0.5)
	}
}

func TestPayloadRegToVar2(t *testing.T) {

	mp := NewMbPayload()
	mp.AddVariable(0, MbVar{name: "T01", fmt: "int16", offset: 0, scale: 10, endian: BIG_ENDIAN})
	mp.AddVariable(1, MbVar{name: "T02", fmt: "uint16", offset: 0, scale: 10, endian: BIG_ENDIAN})
	mp.AddVariable(2, MbVar{name: "T03", fmt: "int32", offset: 0, scale: 10, endian: BIG_ENDIAN})
	mp.AddVariable(4, MbVar{name: "T04", fmt: "uint32", offset: 0, scale: 10, endian: BIG_ENDIAN})
	mp.AddVariable(6, MbVar{name: "T05", fmt: "float32", offset: 0, scale: 1, endian: BIG_ENDIAN})
	vars := mp.regToVar(0, []byte{0, 1, 0, 2, 0, 0, 0, 3, 0, 0, 0, 4, 0x3f, 0, 0, 0})
	if vars[0].Value != 10.0 {
		t.Errorf("got value %f want %f", vars[0].Value, 10.0)
	}
	if vars[1].Value != 20.0 {
		t.Errorf("got value %f want %f", vars[1].Value, 20.0)
	}
	if vars[2].Value != 30.0 {
		t.Errorf("got value %f want %f", vars[2].Value, 30.0)
	}
	if vars[3].Value != 40.0 {
		t.Errorf("got value %f want %f", vars[3].Value, 40.0)
	}
	if vars[4].Value != 0.5 {
		t.Errorf("got value %f want %f", vars[4].Value, 0.5)
	}
}
