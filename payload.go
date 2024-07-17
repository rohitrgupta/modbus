package modbus

type Value struct {
	Name  string
	Value interface{}
}

type MbPayload struct {
	Vars  map[int]MbVar
	Start int
	Size  int
}

func NewMbPayload() MbPayload {
	r := MbPayload{}
	r.Vars = make(map[int]MbVar)
	r.Start = 0
	r.Size = 0
	return r
}

func (r *MbPayload) AddVariable(address int, v MbVar) {
	if len(r.Vars) == 0 {
		r.Start = address
	}
	if address < r.Start {
		r.Size += 2 * (r.Start - address)
		r.Start = address
	}
	r.Vars[address] = v
	var size int
	switch v.Fmt {
	case "uint16", "int16":
		size = 2 * (address - r.Start + 1)
	case "uint32", "int32", "float32":
		size = 2 * (address - r.Start + 2)
	}
	if r.Size < size {
		r.Size = size
	}
}

func (r *MbPayload) RegToVar(start int, reg []byte) []Value {
	if len(reg) < r.Size {
		return nil
	}
	val := make([]Value, len(r.Vars))
	ctr := 0
	for i := 0; i < len(reg); i += 2 {
		address := start + i/2
		if _, ok := r.Vars[address]; ok {
			v := r.Vars[address]
			switch v.Fmt {
			case "uint16", "int16":
				v.SetReg(reg[i : i+2])
				if v.ValueType == VALUE_TYPE_INT {
					val[ctr] = Value{v.Name, v.ValueInt}
				} else {
					val[ctr] = Value{v.Name, v.ValueFloat}
				}
			case "uint32", "int32", "float32":
				v.SetReg(reg[i : i+4])
				if v.ValueType == VALUE_TYPE_INT {
					val[ctr] = Value{v.Name, v.ValueInt}
				} else {
					val[ctr] = Value{v.Name, v.ValueFloat}
				}
				i += 2
			}
			ctr += 1
			r.Vars[address] = v
		}
	}
	return val
}
