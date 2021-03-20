package main

import (
	"./methodarea"
	"fmt"
	"reflect"
)
import "./instruction/base"
import "./instruction"
import "./rt"

func Interpret(method *methodarea.Method) {
	reader := base.NewBytecodeReader(0, method.Code)
	frame := rt.NewStackFrame(rt.NewLocalVarTable(method.MaxLocals), rt.NewOperateStack(method.MaxStack), method)
	Loop(reader, frame)
}

func Loop(reader *base.BytecodeReader, frame *rt.StackFrame) {
	var instr base.Instruction

	for {
		instructionCode := reader.ReadUInt8()

		instr = instruction.NewInstruction(instructionCode)

		_, ok := instr.(*base.NopInstruction)

		fmt.Println(reflect.TypeOf(instr))

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
