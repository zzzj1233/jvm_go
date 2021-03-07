package classfile

type AttributeMarker struct {
}

type AttributeDeprecated struct {
	AttributeMarker
}

type AttributeSynthetic struct {
	AttributeMarker
}

func (this AttributeMarker) readInfo(reader *ClassReader) {

}
