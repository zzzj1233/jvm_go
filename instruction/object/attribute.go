package object

import "../base"
import "../../rt"
import "../../context"
import "../../methodarea"

type GetStatic struct {
	base.Index16Instruction
}

func (this *GetStatic) Execute(frame *rt.StackFrame) {
	invokeClass := frame.Method.Class

	pool := invokeClass.Pool

	fieldRef := pool.GetFieldRef(uint16(this.Index))

	targetClass := context.Loader.LoadClass(fieldRef.GetClassName())

	name, _ := fieldRef.GetNameAndType()

	field := targetClass.AccessStaticField(invokeClass, name)

	if field == nil {
		panic("java/lang/NoSuchFieldError")
	}

	rt.GetAttribute(field, frame)
}

type PutStatic struct {
	base.Index16Instruction
}

func (this *PutStatic) Execute(frame *rt.StackFrame) {
	invokeClass := frame.Method.Class

	pool := invokeClass.Pool

	fieldRef := pool.GetFieldRef(uint16(this.Index))

	targetClass := context.Loader.LoadClass(fieldRef.GetClassName())

	name, _ := fieldRef.GetNameAndType()

	field := targetClass.AccessStaticField(invokeClass, name)

	if field == nil {
		panic("java/lang/NoSuchFieldError")
	}

	rt.PutAttribute(field, frame)
}

type GetField struct {
	base.Index16Instruction
}

func (this *GetField) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	pool := frame.Method.Class.Pool

	accessObject := stack.PopObj()

	fieldRef := pool.GetFieldRef(uint16(this.Index))

	name, _ := fieldRef.GetNameAndType()

	field := accessObject.AccessField(frame.Method.Class, name)

	if field == nil {
		panic("java/lang/NoSuchFieldError")
	}

	rt.GetAttribute(field, frame)
}

type PutField struct {
	base.Index16Instruction
}

func (this *PutField) Execute(frame *rt.StackFrame) {
	stack := frame.OperateStack

	pool := frame.Method.Class.Pool

	fieldRef := pool.GetFieldRef(uint16(this.Index))

	name, desc := fieldRef.GetNameAndType()

	descriptor := methodarea.NewDescriptor(desc)

	var value interface{}

	if descriptor.IsChar || descriptor.IsInt || descriptor.IsShort || descriptor.IsBool || descriptor.IsByte {
		value = stack.PopInt()
	} else if descriptor.IsLong {
		value = stack.PopLong()
	} else if descriptor.IsDouble {
		value = stack.PopDouble()
	} else if descriptor.IsFloat {
		value = stack.PopFloat()
	} else {
		value = stack.PopObj()
	}

	accessObject := stack.PopObj()

	field := accessObject.AccessField(frame.Method.Class, name)

	if field == nil {
		panic("java/lang/NoSuchFieldError")
	}

	field.Value = value
}
