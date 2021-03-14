package condition

import (
	"../base"
)
import "../../rt"

type FCmpL struct {
	base.NoOperandsInstruction
}

func (this *FCmpL) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	float1 := stack.PopFloat()
	float2 := stack.PopFloat()

	if float2 > float1 {
		stack.PushInt(1)
	} else if float2 == float1 {
		stack.PushInt(0)
	} else if float2 < float1 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(-1)
	}

}

type FCmpG struct {
	base.NoOperandsInstruction
}

func (this *FCmpG) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	float1 := stack.PopFloat()
	float2 := stack.PopFloat()

	if float2 > float1 {
		stack.PushInt(1)
	} else if float2 == float1 {
		stack.PushInt(0)
	} else if float2 < float1 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(1)
	}

}
