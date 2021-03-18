package methodarea

type Attribute struct {
	Descriptor *Descriptor
	Value      interface{}
}

func NewAttribute(descriptor string) *Attribute {

	d := newDescriptor(descriptor)

	return &Attribute{
		Descriptor: d,
		Value:      d.DefaultValue(),
	}
}
