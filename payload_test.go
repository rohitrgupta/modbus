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
	mp.AddVariable(104, MbVar{name: "T02", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	if mp.size != size2 {
		t.Errorf("got size %d want %d", mp.size, size2)
	}
	start3 := 90
	size3 := 42
	mp.AddVariable(90, MbVar{name: "T02", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
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
	mp.AddVariable(90, MbVar{name: "T02", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(100, MbVar{name: "T01", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(104, MbVar{name: "T02", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	mp.AddVariable(110, MbVar{name: "T02", fmt: "int16", offset: 0, scale: 1, endian: BIG_ENDIAN})
	if mp.start != start1 {
		t.Errorf("got start %d want %d", mp.start, start1)
	}
	if mp.size != size1 {
		t.Errorf("got size %d want %d", mp.size, size1)
	}

}
