package stack

import "../base"
import "../../rt"

type Pop struct {
	base.NoOperandsInstruction
}

func (this *Pop) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PopSlot()
}

type Pop2 struct {
	base.NoOperandsInstruction
}

func (this *Pop2) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PopSlot()
	frame.OperateStack.PopSlot()
}
