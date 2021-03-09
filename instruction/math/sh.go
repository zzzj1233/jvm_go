package math

// 文档描述

// 通过将value1左移s个位的位置来计算i结果,其中s是value2的低5位的值,结果被压入操作数堆栈。

import "../base"
import "../../rt"

type IShl struct {
	base.NoOperandsInstruction
}

func (this *IShl) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	stack.PushInt(int32(int1) >> (uint32(int2) & 0b00011111))
}

// int右移
type IShR struct {
	base.NoOperandsInstruction
}

func (this *IShR) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	stack.PushInt(int32(int1) >> (uint32(int2) & 0b00011111))
}

// int逻辑右移
type IUShR struct {
	base.NoOperandsInstruction
}

func (this *IUShR) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	stack.PushInt(int32(uint32(int1) >> (uint32(int2) & 0b00011111)))
}

// long左移
type LShl struct {
	base.NoOperandsInstruction
}

func (this *LShl) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	long2 := stack.PopLong()
	int1 := stack.PopInt()
	stack.PushLong(long2 << (uint32(int1) & 0b00011111))
}

// long左移
type LShr struct {
	base.NoOperandsInstruction
}

func (this *LShr) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	long2 := stack.PopLong()
	int1 := stack.PopInt()
	stack.PushLong(int64(long2) >> (uint32(int1) & 0b00111111))
}

// long逻辑右移
type LUShr struct {
	base.NoOperandsInstruction
}

func (this *LUShr) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
