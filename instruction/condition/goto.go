package condition

import "../base"
import "../../rt"

type Goto struct {
	base.BranchInstruction
}

func (this *Goto) Execute(frame *rt.StackFrame) {
	this.MatchCondition(frame)
}
