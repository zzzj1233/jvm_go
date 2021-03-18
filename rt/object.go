package rt

import "../methodarea"

type Object struct {
	class      *methodarea.Class
	Attributes map[int]*methodarea.Attribute
}

func NewObject(class *methodarea.Class) *Object {
	obj := &Object{
		class: class,
	}

	return obj
}

func (this *Object) GetAttribute(index int, frame *StackFrame) {
	m := this.Attributes[index]

	if m == nil {
		panic("no such field error")
	}

	getAttribute(m, frame)
}

func getAttribute(m *methodarea.Attribute, frame *StackFrame) {
	stack := frame.OperateStack
	if m.Descriptor.IsInt ||
		m.Descriptor.IsByte ||
		m.Descriptor.IsChar ||
		m.Descriptor.IsShort ||
		m.Descriptor.IsBool {
		stack.PushInt(m.Value.(int32))
	} else if m.Descriptor.IsLong {
		stack.PushLong(m.Value.(int64))
	} else if m.Descriptor.IsDouble {
		stack.PushDouble(m.Value.(float64))
	} else if m.Descriptor.IsReference {
		stack.PushObj(m.Value.(*Object))
	} else if m.Descriptor.IsArray {
		stack.PushObj(m.Value.(*Object))
	}
}

func (this *Object) GetStaticAttribute(index int, frame *StackFrame) {
	m := this.class.StaticAttributes[index]

	if m == nil {
		panic("no such field error")
	}

	getAttribute(m, frame)
}
