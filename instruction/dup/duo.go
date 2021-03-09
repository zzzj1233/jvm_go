package dup

import "../base"
import "../../rt"

type Dup struct {
	base.NoOperandsInstruction
}

func (this *Dup) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushSlot(frame.OperateStack.TopSlot())
}

type DupX1 struct {
	base.NoOperandsInstruction
}

// 复制一个栈顶变量,并且插入到倒数第一个变量前
func (this *DupX1) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	top := stack.PopSlot()
	slot := stack.PopSlot()

	stack.PushSlot(top)
	stack.PushSlot(slot)
	stack.PushSlot(top)
}

type DupX2 struct {
	base.NoOperandsInstruction
}

// 复制一个栈顶变量,并且插入到倒数第二个变量前
func (this *DupX2) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type Dup2 struct {
	base.NoOperandsInstruction
}

// 复制两个栈顶变量,并且插入到栈顶
func (this *Dup2) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type Dup2X1 struct {
	base.NoOperandsInstruction
}

// 复制两个栈顶变量,并且插入到倒数第一个变量前
func (this *Dup2X1) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type Dup2X2 struct {
	base.NoOperandsInstruction
}

// 复制两个栈顶变量,并且插入到倒数第二个变量前
func (this *Dup2X2) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
