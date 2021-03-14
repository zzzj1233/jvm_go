package math

// 异或

import "../base"
import "../../rt"

type IXor struct {
	base.NoOperandsInstruction
}

func (this *IXor) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	stack.PushInt(int1 ^ int2)
}

type LXor struct {
	base.NoOperandsInstruction
}

func (this *LXor) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	long2 := stack.PopLong()
	long1 := stack.PopLong()
	stack.PushLong(long1 ^ long2)
}
