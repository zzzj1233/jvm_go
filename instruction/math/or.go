package math

// 按位或
import "../base"
import "../../rt"

type IOr struct {
	base.NopInstruction
}

func (this *IOr) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	stack.PushInt(int1 | int2)
}

type LOr struct {
	base.NopInstruction
}

func (this *LOr) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	long2 := stack.PopLong()
	long1 := stack.PopLong()
	stack.PushLong(long2 | long1)
}
