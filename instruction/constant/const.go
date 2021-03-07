package constant

import "../base"
import "../../rt"

type AConstNull struct {
	base.NoOperandsInstruction
}

// 推入一个null到操作数栈
func (this *AConstNull) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushObj(nil)
}

type IConstM1 struct {
	base.NoOperandsInstruction
}

// 推入-1到操作数栈
func (this *IConstM1) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(-1)
}

type IConst0 struct {
	base.NoOperandsInstruction
}

// 推入0到操作数栈
func (this *IConst0) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(0)
}

type IConst1 struct {
	base.NoOperandsInstruction
}

func (this *IConst1) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(1)
}

type IConst2 struct {
	base.NoOperandsInstruction
}

func (this *IConst2) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(2)
}

type IConst3 struct {
	base.NoOperandsInstruction
}

func (this *IConst3) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(3)
}

type IConst4 struct {
	base.NoOperandsInstruction
}

func (this *IConst4) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(4)
}

type IConst5 struct {
	base.NoOperandsInstruction
}

func (this *IConst5) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(5)
}

type FConst0 struct {
	base.NoOperandsInstruction
}

func (this *FConst0) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushFloat(0)
}

type FConst1 struct {
	base.NoOperandsInstruction
}

func (this *FConst1) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushFloat(1)
}

type FConst2 struct {
	base.NoOperandsInstruction
}

func (this *FConst2) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushFloat(2)
}

type DConst0 struct {
	base.NoOperandsInstruction
}

func (this *DConst0) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushDouble(0)
}

type DConst1 struct {
	base.NoOperandsInstruction
}

func (this *DConst1) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushDouble(1)
}

type LConst0 struct {
	base.NoOperandsInstruction
}

func (this *LConst0) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushLong(0)
}

type LConst1 struct {
	base.NoOperandsInstruction
}

func (this *LConst1) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushLong(1)
}
