package base

import "../../rt"

type BranchInstruction struct {
	Offset int16
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset = reader.ReadInt16()
}

func (this *BranchInstruction) MatchCondition(frame *rt.StackFrame) {
	// 减去一个指令 + 两个操作数的长度
	frame.NextPc = uint(this.Offset) - 3
}
