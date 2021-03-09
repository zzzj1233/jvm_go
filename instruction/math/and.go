package math

import "../base"
import "../../rt"

type IAnd struct {
	base.NopInstruction
}

func (this *IAnd) Execute(frame *rt.StackFrame) {
	int2 := frame.OperateStack.PopInt()
	int1 := frame.OperateStack.PopInt()

	result := int1 & int2
	frame.OperateStack.PushInt(result)
}

type LAnd struct {
	base.NopInstruction
}

func (this *LAnd) Execute(frame *rt.StackFrame) {
	long2 := frame.OperateStack.PopLong()
	long1 := frame.OperateStack.PopLong()

	result := long1 & long2
	frame.OperateStack.PushLong(result)
}
