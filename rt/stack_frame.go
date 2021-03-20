package rt

import "../methodarea"

type StackFrame struct {
	LocalVarTable LocalVarTable
	OperateStack  *OperateStack
	Low           *StackFrame
	NextPc        int
	Method        *methodarea.Method
}

func NewStackFrame(localVarTable LocalVarTable, operateStack *OperateStack, method *methodarea.Method) *StackFrame {
	return &StackFrame{
		LocalVarTable: localVarTable,
		OperateStack:  operateStack,
		Method:        method,
	}
}
