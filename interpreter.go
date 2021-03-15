package main

import (
	"./classfile"
	"fmt"
)
import "./instruction/base"
import "./instruction"
import "./rt"

func Interpret(method *classfile.MemberInfo) {
	for _, attr := range method.Attributes {
		codeAttribute := attr.(*classfile.CodeAttribute)
		reader := base.NewBytecodeReader(0, codeAttribute.GetCode())
		frame := rt.NewStackFrame(rt.NewLocalVarTable(codeAttribute.MaxLocals), rt.NewOperateStack(codeAttribute.MaxStack))
		Loop(reader, frame)
	}
}

func Loop(reader *base.BytecodeReader, frame *rt.StackFrame) {
	var instr base.Instruction

	for {
		instructionCode := reader.ReadUInt8()

		instr = instruction.NewInstruction(instructionCode)

		_, ok := instr.(*base.NopInstruction)

		if ok {
			fmt.Println("table = ", frame.LocalVarTable)
			fmt.Println("stack = ", frame.OperateStack)
			return
		}

		instr.FetchOperands(reader)

		instr.Execute(frame)

		if frame.NextPc != 0 {
			reader.SetNextPc(uint(frame.NextPc))
			frame.NextPc = 0
		}
	}

}
