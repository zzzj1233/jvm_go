package math

import "../base"
import "../../rt"

type DMul struct {
	base.NopInstruction
}

func (this *DMul) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushDouble(frame.OperateStack.PopDouble() * frame.OperateStack.PopDouble())
}

type FMul struct {
	base.NopInstruction
}

func (this *FMul) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushFloat(frame.OperateStack.PopFloat() * frame.OperateStack.PopFloat())
}

type IMul struct {
	base.NopInstruction
}

func (this *IMul) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(frame.OperateStack.PopInt() * frame.OperateStack.PopInt())
}

type LMul struct {
	base.NopInstruction
}

func (this *LMul) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushLong(frame.OperateStack.PopLong() * frame.OperateStack.PopLong())
}
