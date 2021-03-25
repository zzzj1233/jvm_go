package base

import "../../rt"
import "../../methodarea"

type Instruction interface {
	// 获取操作数,某些指令无需获取操作数
	FetchOperands(reader *BytecodeReader)

	Execute(frame *rt.StackFrame)
}

type NopInstruction struct {
	NoOperandsInstruction
}

func (this *NopInstruction) Execute(frame *rt.StackFrame) {
	// no op
}

type NoOperandsInstruction struct {
}

// 不需要读取操作数
func (this *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// no op
}

type Index8Instruction struct {
	Index int
}

func (this *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = int(reader.ReadInt8())
}

type Index16Instruction struct {
	Index int16
}

func (this *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = int16(reader.ReadUInt8())<<8 | int16(reader.ReadUInt8())
}

func InvokeMethod(method *methodarea.Method, frame *rt.StackFrame) {
	newFrame := rt.NewStackFrame(rt.NewLocalVarTable(method.MaxLocals), rt.NewOperateStack(method.MaxStack), method, frame.Thread)

	for i := 0; i < method.ArgSlotCount(); i++ {
		newFrame.LocalVarTable.SetSlot(i, frame.OperateStack.PopSlot())
	}

	frame.Thread.Stack.PushFrame(newFrame)
}
