package object

import "../base"
import "../../rt"

type InvokeSpecial struct {
	base.Index16Instruction
}

func (this *InvokeSpecial) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PopObj()
}
