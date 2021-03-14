package cast

import "../base"
import "../../rt"

type I2B struct {
	base.NoOperandsInstruction
}

func (this *I2B) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushInt(int32(int8(stack.PopInt())))
}

type I2C struct {
	base.NoOperandsInstruction
}

func (this *I2C) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushInt(int32(uint16(stack.PopInt())))
}

type I2D struct {
	base.NoOperandsInstruction
}

func (this *I2D) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushDouble(float64(stack.PopInt()))
}

type I2F struct {
	base.NoOperandsInstruction
}

func (this *I2F) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushFloat(float32(stack.PopInt()))
}

type I2L struct {
	base.NoOperandsInstruction
}

func (this *I2L) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushLong(int64(stack.PopInt()))
}

type I2S struct {
	base.NoOperandsInstruction
}

func (this *I2S) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	stack.PushInt(int32(int16(stack.PopInt())))
}
