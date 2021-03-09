package math

import "../base"
import "../../rt"

type DAdd struct {
	base.NoOperandsInstruction
}

func (this *DAdd) Execute(frame *rt.StackFrame) {
	double1 := frame.OperateStack.PopDouble()
	double2 := frame.OperateStack.PopDouble()

	result := double1 + double2
	frame.OperateStack.PushDouble(result)
}

type FAdd struct {
	base.NoOperandsInstruction
}

func (this *FAdd) Execute(frame *rt.StackFrame) {
	float1 := frame.OperateStack.PopFloat()
	float2 := frame.OperateStack.PopFloat()

	result := float1 + float2
	frame.OperateStack.PushFloat(result)
}

type IAdd struct {
	base.NoOperandsInstruction
}

func (this *IAdd) Execute(frame *rt.StackFrame) {
	int1 := frame.OperateStack.PopInt()
	int2 := frame.OperateStack.PopInt()

	result := int1 + int2
	frame.OperateStack.PushInt(result)
}

type LAdd struct {
	base.NoOperandsInstruction
}

func (this *LAdd) Execute(frame *rt.StackFrame) {
	long1 := frame.OperateStack.PopLong()
	long2 := frame.OperateStack.PopLong()

	result := long1 + long2
	frame.OperateStack.PushLong(result)
}
