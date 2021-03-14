package condition

import "../base"
import "../../rt"

type AcmpEq struct {
	base.BranchInstruction
}

func (this *AcmpEq) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	obj1 := stack.PopObj()
	obj2 := stack.PopObj()

	if obj1 == obj2 {
		this.MatchCondition(frame)
	}
}

type AcmpNe struct {
	base.BranchInstruction
}

func (this *AcmpNe) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	obj1 := stack.PopObj()
	obj2 := stack.PopObj()

	if obj1 != obj2 {
		this.MatchCondition(frame)
	}
}

type IfNonNull struct {
	base.BranchInstruction
}

func (this *IfNonNull) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	obj1 := stack.PopObj()

	if obj1 != nil {
		this.MatchCondition(frame)
	}
}

type IfNull struct {
	base.BranchInstruction
}

func (this *IfNull) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	obj1 := stack.PopObj()

	if obj1 == nil {
		this.MatchCondition(frame)
	}
}
