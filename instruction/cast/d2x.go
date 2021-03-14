package cast

import "../base"
import "../../rt"

type D2F struct {
	base.NoOperandsInstruction
}

func (this *D2F) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	double := stack.PopDouble()
	stack.PushFloat(float32(double))
}

type D2I struct {
	base.NoOperandsInstruction
}

func (this *D2I) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	double := stack.PopDouble()
	stack.PushInt(int32(double))
}

type D2L struct {
	base.NoOperandsInstruction
}

func (this *D2L) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	double := stack.PopDouble()
	stack.PushLong(int64(double))
}
