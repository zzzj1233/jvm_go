package cast

import "../base"
import "../../rt"

type L2F struct {
	base.NoOperandsInstruction
}

func (this *L2F) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushFloat(float32(stack.PopLong()))
}

type L2D struct {
	base.NoOperandsInstruction
}

func (this *L2D) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushDouble(float64(stack.PopLong()))
}

type L2I struct {
	base.NoOperandsInstruction
}

func (this *L2I) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushInt(int32(stack.PopLong()))
}
