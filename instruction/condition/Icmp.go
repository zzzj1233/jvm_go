package condition

import "../base"
import "../../rt"

type IcmpEq struct {
	base.BranchInstruction
}

func (this *IcmpEq) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()
	int2 := stack.PopInt()

	if int1 == int2 {
		this.MatchCondition(frame)
	}

}

type IcmpNe struct {
	base.BranchInstruction
}

func (this *IcmpNe) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()
	int2 := stack.PopInt()

	if int1 != int2 {
		this.MatchCondition(frame)
	}

}

type IcmpLt struct {
	base.BranchInstruction
}

func (this *IcmpLt) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()
	int2 := stack.PopInt()

	if int1 < int2 {
		this.MatchCondition(frame)
	}

}

type IcmpLe struct {
	base.BranchInstruction
}

func (this *IcmpLe) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()
	int2 := stack.PopInt()

	if int1 <= int2 {
		this.MatchCondition(frame)
	}

}

type IcmpGt struct {
	base.BranchInstruction
}

func (this *IcmpGt) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()
	int2 := stack.PopInt()

	if int1 > int2 {
		this.MatchCondition(frame)
	}

}

type IcmpGe struct {
	base.BranchInstruction
}

func (this *IcmpGe) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()
	int2 := stack.PopInt()

	if int1 >= int2 {
		this.MatchCondition(frame)
	}

}

type IfEq struct {
	base.BranchInstruction
}

func (this *IfEq) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()

	if int1 == 0 {
		this.MatchCondition(frame)
	}

}

type IfNe struct {
	base.BranchInstruction
}

func (this *IfNe) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()

	if int1 != 0 {
		this.MatchCondition(frame)
	}

}

type IfLt struct {
	base.BranchInstruction
}

func (this *IfLt) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()

	if int1 < 0 {
		this.MatchCondition(frame)
	}

}

type IfLe struct {
	base.BranchInstruction
}

func (this *IfLe) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()

	if int1 <= 0 {
		this.MatchCondition(frame)
	}

}

type IfGt struct {
	base.BranchInstruction
}

func (this *IfGt) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()

	if int1 > 0 {
		this.MatchCondition(frame)
	}

}

type IfGe struct {
	base.BranchInstruction
}

func (this *IfGe) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack
	int1 := stack.PopInt()

	if int1 >= 0 {
		this.MatchCondition(frame)
	}

}
