package math

// 取反
import "../base"
import "../../rt"

type DNeg struct {
	base.NoOperandsInstruction
}

func (this *DNeg) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushDouble(-frame.OperateStack.PopDouble())
}

type FNeg struct {
	base.NoOperandsInstruction
}

func (this *FNeg) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushFloat(-frame.OperateStack.PopFloat())
}

type INeg struct {
	base.NoOperandsInstruction
}

func (this *INeg) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(-frame.OperateStack.PopInt())
}

type LNeg struct {
	base.NoOperandsInstruction
}

func (this *LNeg) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushLong(-frame.OperateStack.PopLong())
}
