package methodarea

import "fmt"

type Attribute struct {
	Field      *Field
	Descriptor *Descriptor
	Value      interface{}
}

func (this *Attribute) String() string {
	return "desc = " + this.Descriptor.desc + ", value = " + fmt.Sprintf("%v", this.Value)
}

func NewAttribute(field *Field) *Attribute {

	d := NewDescriptor(field.Descriptor)

	return &Attribute{
		Descriptor: d,
		Value:      d.DefaultValue(),
		Field:      field,
	}
}
