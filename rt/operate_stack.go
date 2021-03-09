package rt

import (
	"math"
)

type OperateStack struct {
	slots []Slot
	size  uint
}

func NewOperateStack(maxSize uint16) *OperateStack {
	return &OperateStack{
		slots: make([]Slot, maxSize),
		size:  0,
	}
}

func (this *OperateStack) PushInt(val int32) {
	this.slots[this.size].number = val
	this.size++
}

func (this *OperateStack) PopInt() int32 {
	this.size--
	return this.slots[this.size].number
}

func (this *OperateStack) PushFloat(val float32) {
	this.slots[this.size].number = int32(math.Float32bits(val))
	this.size++
}

func (this *OperateStack) PopFloat() float32 {
	this.size--
	return math.Float32frombits(uint32(this.slots[this.size].number))
}

func (this *OperateStack) PushLong(val int64) {
	this.slots[this.size].number = int32(val)
	this.slots[this.size+1].number = int32(val >> 32)
	this.size += 2
}

func (this *OperateStack) PopLong() int64 {
	this.size--
	high := int64(this.slots[this.size].number) << 32
	this.size--
	low := int64(this.slots[this.size].number)

	return high | low
}

func (this *OperateStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	this.PushLong(int64(bits))
}

func (this *OperateStack) PopDouble() float64 {
	long := this.PopLong()
	return math.Float64frombits(uint64(long))
}

func (this *OperateStack) PushObj(obj *Object) {
	this.slots[this.size].obj = obj
	this.size++
}

func (this *OperateStack) PopObj() *Object {
	obj := this.slots[this.size].obj
	this.size--
	this.slots[this.size].obj = nil
	return obj
}

func (this *OperateStack) TopSlot() Slot {
	return this.slots[this.size-1]
}

func (this *OperateStack) PushSlot(slot Slot) {
	this.slots[this.size] = slot
	this.size++
}

func (this *OperateStack) PopSlot() Slot {
	this.size--
	return this.slots[this.size]
}
