package method

import "../base"
import "../../rt"
import "../../methodarea"
import "../../classfile"
import "../../context"

type InvokeSpecial struct {
	base.Index16Instruction
}

func (this *InvokeSpecial) Execute(frame *rt.StackFrame) {
	pool := frame.Method.Pool

	methodRef := pool.GetMethodRef(uint16(this.Index))

	method := loopUpMethod(methodRef)

	base.InvokeMethod(method, frame)
}

func loopUpMethod(methodRef *classfile.ConstantMethodRefInfo) *methodarea.Method {

	methodClass := context.Loader.LoadClass(methodRef.GetClassName())

	methodName, _ := methodRef.GetNameAndType()

	methods := methodClass.GetMethods()

	if methodName != "<init>" {
		panic("java.lang.NoSuchMethodError : " + methodName)
	}

	for _, method := range methods {
		if method.Name == methodName {
			// 是否有访问权限
			// initMethod只允许当前类调用,及当前类的子类调用
			if !methodClass.IsAssignableFrom(method.Class) {
				panic("java.lang.NoSuchMethodError : " + methodName)
			}
			return method
		}
	}

	panic("java.lang.NoSuchMethodError : " + methodName)
	return nil
}
