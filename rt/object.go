package rt

import "../methodarea"

type ObjectAttribute struct {
	field *methodarea.Field
	value interface{}
}

func NewObjectAttribute(field *methodarea.Field) *ObjectAttribute {
	descriptor := field.FieldDescriptor

	return &ObjectAttribute{
		field: field,
		value: descriptor.DefaultValue(),
	}
}

type Object struct {
}
