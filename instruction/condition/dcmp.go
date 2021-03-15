package condition

import (
	"../base"
)
import "../../rt"

type DCmpL struct {
	base.NoOperandsInstruction
}

func (this *DCmpL) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	double1 := stack.PopDouble()
	double2 := stack.PopDouble()

	if double2 > double1 {
		stack.PushInt(1)
	} else if double2 == double1 {
		stack.PushInt(0)
	} else if double1 < double1 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(-1)
	}

}

type DCmpG struct {
	base.NoOperandsInstruction
}

func (this *DCmpG) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	double1 := stack.PopDouble()
	double2 := stack.PopDouble()

	if double2 > double1 {
		stack.PushInt(1)
	} else if double2 == double1 {
		stack.PushInt(0)
	} else if double1 < double1 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(1)
	}

}
