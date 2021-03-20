package object

import "../base"
import "../../rt"
import "../../context"

type Instanceof struct {
	base.Index16Instruction
}

func (this *Instanceof) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	obj := stack.PopObj()

	classRef := frame.Method.Class.Pool.GetClassInfo(uint16(this.Index))

	class := context.Loader.LoadClass(classRef.GetClassName())

	if obj != nil && obj.IsAssignableFrom(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}

}

type CheckCast struct {
	base.Index16Instruction
}

func (this *CheckCast) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	obj := stack.PopObj()

	classRef := frame.Method.Class.Pool.GetClassInfo(uint16(this.Index))

	class := context.Loader.LoadClass(classRef.GetClassName())

	if obj == nil || !obj.IsAssignableFrom(class) {
		panic("java.lang.ClassCastException")
	}

}
