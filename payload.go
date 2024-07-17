package modbus

type Value struct {
	Name  string
	Value interface{}
}

type MbPayload struct {
	vars  map[int]MbVar
	start int
	size  int
}

func NewMbPayload() MbPayload {
	r := MbPayload{}
	r.vars = make(map[int]MbVar)
	r.start = 0
	r.size = 0
	return r
}

func (r *MbPayload) AddVariable(address int, v MbVar) {
	if len(r.vars) == 0 {
		r.start = address
	}
	if address < r.start {
		r.size += 2 * (r.start - address)
		r.start = address
	}
	r.vars[address] = v
	var size int
	switch v.Fmt {
	case "uint16", "int16":
		size = 2 * (address - r.start + 1)
	case "uint32", "int32", "float32":
		size = 2 * (address - r.start + 2)
	}
	if r.size < size {
		r.size = size
	}
}

func (r *MbPayload) regToVar(start int, reg []byte) []Value {
	if len(reg) < r.size {
		return nil
	}
	val := make([]Value, len(r.vars))
	ctr := 0
	for i := 0; i < len(reg); i += 2 {
		address := start + i/2
		if _, ok := r.vars[address]; ok {
			v := r.vars[address]
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
			r.vars[address] = v
		}
	}
	return val
}
