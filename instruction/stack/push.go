package stack

import "../base"
import "../../rt"

type BiPush struct {
	val int8
}

// 抓取一个byte推入到操作数栈
func (this *BiPush) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt8()
}

func (this *BiPush) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(int32(this.val))
}

type SiPush struct {
	val int16
}

// 将一个short推入到操作数栈
func (this *SiPush) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt16()
}

func (this *SiPush) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(int32(this.val))
}
