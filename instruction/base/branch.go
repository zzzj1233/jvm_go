package base

import "../../rt"

type BranchInstruction struct {
	Offset uint16
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset = reader.ReadUInt16()
}

func (this *BranchInstruction) MatchCondition(frame *rt.StackFrame) {
	frame.NextPc = int(this.Offset)
}
