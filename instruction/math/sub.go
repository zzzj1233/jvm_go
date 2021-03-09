package math

import "../base"
import "../../rt"

type DSub struct {
	base.NoOperandsInstruction
}

func (this *DSub) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	double2 := stack.PopDouble()
	double1 := stack.PopDouble()
	stack.PushDouble(double1 - double2)
}

type FSub struct {
	base.NoOperandsInstruction
}

func (this *FSub) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	float2 := stack.PopFloat()
	float1 := stack.PopFloat()
	stack.PushFloat(float1 - float2)
}

type ISub struct {
	base.NoOperandsInstruction
}

func (this *ISub) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	stack.PushInt(int1 - int2)
}

type LSub struct {
	base.NoOperandsInstruction
}

func (this *LSub) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	long2 := stack.PopLong()
	long1 := stack.PopLong()
	stack.PushLong(long1 - long2)
}
