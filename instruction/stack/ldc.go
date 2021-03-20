package stack

import "../base"
import "../../rt"
import "../../classfile"

type Ldc struct {
	base.Index8Instruction
}

type LdcW struct {
	base.Index16Instruction
}

type Ldc2W struct {
	base.Index16Instruction
}

func _execute(index int, frame *rt.StackFrame) {
	pool := frame.Method.Pool

	info := pool[index]

	stack := frame.OperateStack

	switch info.(type) {
	case *classfile.ConstantIntegerInfo:
		stack.PushInt(info.(*classfile.ConstantIntegerInfo).Value)
	case *classfile.ConstantFloatInfo:
		stack.PushFloat(info.(*classfile.ConstantFloatInfo).Value)
	default:
		panic("todo ldc")
	}
}

func (this *Ldc) Execute(frame *rt.StackFrame) {
	_execute(this.Index-1, frame)
}

func (this *LdcW) Execute(frame *rt.StackFrame) {
	_execute(int(this.Index)-1, frame)
}

func (this *Ldc2W) Execute(frame *rt.StackFrame) {
	pool := frame.Method.Pool

	info := pool[this.Index-1]

	stack := frame.OperateStack

	switch info.(type) {
	case *classfile.ConstantLongInfo:
		stack.PushLong(info.(*classfile.ConstantLongInfo).Value)
	case *classfile.ConstantDoubleInfo:
		stack.PushDouble(info.(*classfile.ConstantDoubleInfo).Value)
	default:
		panic("todo ldc2w")
	}
}
