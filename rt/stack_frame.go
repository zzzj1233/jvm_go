package rt

type StackFrame struct {
	LocalVarTable LocalVarTable
	OperateStack  *OperateStack
	Low           *StackFrame
	NextPc        int
}

func NewStackFrame(localVarTable LocalVarTable, operateStack *OperateStack) *StackFrame {
	return &StackFrame{
		LocalVarTable: localVarTable,
		OperateStack:  operateStack,
	}
}
