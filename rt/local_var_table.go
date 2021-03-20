package rt

import "math"

type LocalVarTable []Slot

type Slot struct {
	number int32
	obj    *Object
}

func (this *Slot) GetNumber() int32 {
	return this.number
}

func (this *Slot) GetObj() *Object {
	return this.obj
}

func NewLocalVarTable(size uint16) LocalVarTable {
	return make([]Slot, size)
}

func (this LocalVarTable) PutInt(index int, val int32) {
	this[index].number = val
}

func (this LocalVarTable) GetInt(index int) int32 {
	return this[index].number
}

func (this LocalVarTable) PutFloat(index int, val float32) {
	bits := math.Float32bits(val)
	this[index].number = int32(bits)
}

func (this LocalVarTable) GetFloat(index int) float32 {
	return math.Float32frombits(uint32(this[index].number))
}

func (this LocalVarTable) PutLong(index int, val int64) {
	this[index].number = int32(val)
	this[index+1].number = int32(val >> 32)
}

func (this LocalVarTable) GetLong(index int) int64 {
	return int64(this[index+1].number)<<32 | int64(this[index].number)
}

func (this LocalVarTable) PutDouble(index int, val float64) {
	bits := math.Float64bits(val)
	this.PutLong(index, int64(bits))
}

func (this LocalVarTable) GetDouble(index int) float64 {
	bits := uint64(this.GetLong(index))
	return math.Float64frombits(bits)
}

func (this LocalVarTable) PutRef(index int, object *Object) {
	this[index].obj = object
}

func (this LocalVarTable) GetRef(index int) *Object {
	return this[index].obj
}
