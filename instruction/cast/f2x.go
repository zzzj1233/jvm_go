package cast

import "../base"
import "../../rt"

type F2D struct {
	base.NoOperandsInstruction
}

func (this *F2D) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushDouble(float64(stack.PopFloat()))
}

type F2I struct {
	base.NoOperandsInstruction
}

func (this *F2I) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushInt(int32(stack.PopFloat()))
}

type F2L struct {
	base.NoOperandsInstruction
}

func (this *F2L) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushLong(int64(stack.PopFloat()))
}
