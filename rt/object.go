package rt

import (
	"../methodarea"
)

type Object struct {
	Class      *methodarea.Class
	Attributes map[string]*methodarea.Attribute
}

func (this *Object) String() string {
	return "Class {" + this.Class.GetClassName() + " }"
}

func (this *Object) AccessField(visitor *methodarea.Class, name string) *methodarea.Attribute {
	attribute := this.Attributes[name]

	if attribute == nil {
		panic("java.lang.NoSuchFieldError")
	}

	this.Class.Access(visitor, attribute.Field.Flag.ClassMemberFlag)

	return attribute
}
func (obj *Object) IsAssignableFrom(other *methodarea.Class) bool {
	return obj.Class.IsAssignableFrom(other)
}

func NewObject(class *methodarea.Class) *Object {

	size := len(class.InstanceFieldMap)

	attrs := make(map[string]*methodarea.Attribute, size)

	for k, v := range class.InstanceFieldMap {
		attrs[k] = methodarea.NewAttribute(v)
	}

	obj := &Object{
		Class:      class,
		Attributes: attrs,
	}

	return obj
}

func PutAttribute(m *methodarea.Attribute, frame *StackFrame) {
	stack := frame.OperateStack
	if m.Descriptor.IsInt ||
		m.Descriptor.IsByte ||
		m.Descriptor.IsChar ||
		m.Descriptor.IsShort ||
		m.Descriptor.IsBool {
		m.Value = stack.PopInt()
	} else if m.Descriptor.IsLong {
		m.Value = stack.PopLong()
	} else if m.Descriptor.IsDouble {
		m.Value = stack.PopDouble()
	} else if m.Descriptor.IsReference {
		m.Value = stack.PopObj()
	} else if m.Descriptor.IsArray {
		m.Value = stack.PopObj()
	}
}

func GetAttribute(m *methodarea.Attribute, frame *StackFrame) {
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
		if m.Value == nil {
			stack.PushObj(nil)
		} else {
			stack.PushObj(m.Value.(*Object))
		}
	} else if m.Descriptor.IsArray {
		if m.Value == nil {
			stack.PushObj(nil)
		} else {
			stack.PushObj(m.Value.(*Object))
		}
	}
}
