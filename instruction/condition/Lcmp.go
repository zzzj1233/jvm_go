package condition

import (
	"../base"
)
import "../../rt"

type LCmp struct {
	base.NoOperandsInstruction
}

func (this *LCmp) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	long1 := stack.PopLong()
	long2 := stack.PopLong()

	if long2 > long1 {
		stack.PushInt(1)
	} else if long2 == long1 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}

}
