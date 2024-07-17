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
			mp.AddVariable(tt.start, MbVar{Name: tt.name, Fmt: tt.fmt, Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
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
	mp.AddVariable(start1, MbVar{Name: "T01", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	if mp.start != start1 {
		t.Errorf("got start %d want %d", mp.start, start1)
	}
	if mp.size != size1 {
		t.Errorf("got size %d want %d", mp.size, size1)
	}
	size2 := 22
	mp.AddVariable(110, MbVar{Name: "T02", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	if mp.start != start1 {
		t.Errorf("got start %d want %d", mp.start, start1)
	}
	if mp.size != size2 {
		t.Errorf("got size %d want %d", mp.size, size2)
	}
	mp.AddVariable(104, MbVar{Name: "T03", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	if mp.size != size2 {
		t.Errorf("got size %d want %d", mp.size, size2)
	}
	start3 := 90
	size3 := 42
	mp.AddVariable(90, MbVar{Name: "T04", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
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
	mp.AddVariable(90, MbVar{Name: "T01", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	mp.AddVariable(100, MbVar{Name: "T02", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	mp.AddVariable(104, MbVar{Name: "T03", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	mp.AddVariable(110, MbVar{Name: "T04", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	if mp.start != start1 {
		t.Errorf("got start %d want %d", mp.start, start1)
	}
	if mp.size != size1 {
		t.Errorf("got size %d want %d", mp.size, size1)
	}

}

func TestPayloadRegToVar1(t *testing.T) {

	mp := NewMbPayload()
	mp.AddVariable(0, MbVar{Name: "T01", Fmt: "int16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	mp.AddVariable(1, MbVar{Name: "T02", Fmt: "uint16", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	mp.AddVariable(2, MbVar{Name: "T03", Fmt: "int32", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	mp.AddVariable(4, MbVar{Name: "T04", Fmt: "uint32", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	mp.AddVariable(6, MbVar{Name: "T05", Fmt: "float32", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
	vars := mp.regToVar(0, []byte{0, 1, 0, 2, 0, 0, 0, 3, 0, 0, 0, 4, 0x3f, 0, 0, 0})
	if mp.vars[0].ValueInt != 1 {
		t.Errorf("got value %d want %d", mp.vars[0].ValueInt, 1)
	}
	if mp.vars[1].ValueInt != 2 {
		t.Errorf("got value %d want %d", mp.vars[1].ValueInt, 2)
	}
	if mp.vars[2].ValueInt != 3 {
		t.Errorf("got value %d want %d", mp.vars[2].ValueInt, 3)
	}
	if mp.vars[4].ValueInt != 4 {
		t.Errorf("got value %d want %d", mp.vars[3].ValueInt, 4)
	}
	if mp.vars[6].ValueFloat != 0.5 {
		t.Errorf("got value %f want %f", mp.vars[3].ValueFloat, 0.5)
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
	mp.AddVariable(0, MbVar{Name: "T01", Fmt: "int16", Offset: 0, Scale: 10, Endian: BIG_ENDIAN})
	mp.AddVariable(1, MbVar{Name: "T02", Fmt: "uint16", Offset: 0, Scale: 10, Endian: BIG_ENDIAN})
	mp.AddVariable(2, MbVar{Name: "T03", Fmt: "int32", Offset: 0, Scale: 10, Endian: BIG_ENDIAN})
	mp.AddVariable(4, MbVar{Name: "T04", Fmt: "uint32", Offset: 0, Scale: 10, Endian: BIG_ENDIAN})
	mp.AddVariable(6, MbVar{Name: "T05", Fmt: "float32", Offset: 0, Scale: 1, Endian: BIG_ENDIAN})
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
