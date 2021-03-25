package main

import (
	"./methodarea"
	"fmt"
	"reflect"
)
import "./instruction/base"
import "./instruction"
import "./rt"

func Interpret(method *methodarea.Method, thread *rt.Thread) {
	reader := base.NewBytecodeReader(0, method.Code)
	frame := rt.NewStackFrame(rt.NewLocalVarTable(method.MaxLocals), rt.NewOperateStack(method.MaxStack), method, thread)
	thread.Stack.PushFrame(frame)
	Loop(reader, thread)
}

func Loop(reader *base.BytecodeReader, thread *rt.Thread) {
	var instr base.Instruction

	for {
		frame := thread.Stack.TopFrame()

		reader.Reset(frame.NextPc, frame.Method.Code)

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

		frame.NextPc = reader.GetPc()

		instr.Execute(frame)
	}

}
