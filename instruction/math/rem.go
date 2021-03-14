package math

import (
	"../base"
	"math"
)
import "../../rt"

type DRem struct {
	base.NoOperandsInstruction
}

func (this *DRem) Execute(frame *rt.StackFrame) {
	double2 := frame.OperateStack.PopDouble()
	double1 := frame.OperateStack.PopDouble()

	frame.OperateStack.PushDouble(math.Mod(double1, double2))
}

type FRem struct {
	base.NoOperandsInstruction
}

func (this *FRem) Execute(frame *rt.StackFrame) {
	float2 := frame.OperateStack.PopFloat()
	float1 := frame.OperateStack.PopFloat()

	frame.OperateStack.PushFloat(float32(math.Mod(float64(float1), float64(float2))))
}

type IRem struct {
	base.NoOperandsInstruction
}

func (this *IRem) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	int2 := stack.PopInt()
	int1 := stack.PopInt()

	stack.PushInt(int1 / int2)
}

type LRem struct {
	base.NoOperandsInstruction
}

func (this *LRem) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	long2 := stack.PopLong()
	long1 := stack.PopLong()

	stack.PushLong(long2 / long1)
}
