package object

import "../base"
import "../../rt"
import "../../context"

type New struct {
	base.Index16Instruction
}

func (this *New) Execute(frame *rt.StackFrame) {
	classRef := frame.Method.Pool.GetClassInfo(uint16(this.Index))

	if classRef == nil {
		panic("java.lang.ClassNotFoundError")
	}

	class := context.Loader.LoadClass(classRef.GetClassName())

	object := rt.NewObject(class)

	frame.OperateStack.PushObj(object)
}
