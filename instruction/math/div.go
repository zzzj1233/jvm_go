package math

import "../base"
import "../../rt"

type DDiv struct {
	base.NopInstruction
}

func (this *DDiv) Execute(frame *rt.StackFrame) {
	double2 := frame.OperateStack.PopDouble()
	double1 := frame.OperateStack.PopDouble()

	frame.OperateStack.PushDouble(double1 / double2)
}

type FDiv struct {
	base.NopInstruction
}

func (this *FDiv) Execute(frame *rt.StackFrame) {
	float2 := frame.OperateStack.PopFloat()
	float1 := frame.OperateStack.PopFloat()

	frame.OperateStack.PushFloat(float1 / float2)
}

type IDiv struct {
	base.NopInstruction
}

func (this *IDiv) Execute(frame *rt.StackFrame) {
	int2 := frame.OperateStack.PopInt()
	int1 := frame.OperateStack.PopInt()

	frame.OperateStack.PushInt(int1 / int2)
}

type LDiv struct {
	base.NopInstruction
}

func (this *LDiv) Execute(frame *rt.StackFrame) {
	long2 := frame.OperateStack.PopLong()
	long1 := frame.OperateStack.PopLong()

	frame.OperateStack.PushLong(long1 / long2)
}
