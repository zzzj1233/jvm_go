package rt

import "../methodarea"

type StackFrame struct {
	LocalVarTable LocalVarTable
	OperateStack  *OperateStack
	Low           *StackFrame
	NextPc        uint
	Method        *methodarea.Method
	Thread        *Thread
}

func NewStackFrame(localVarTable LocalVarTable, operateStack *OperateStack, method *methodarea.Method, thread *Thread) *StackFrame {
	return &StackFrame{
		LocalVarTable: localVarTable,
		OperateStack:  operateStack,
		Method:        method,
		Thread:        thread,
	}
}
