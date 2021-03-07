package classfile

type AttributeUnParse struct {
	attrLen uint32
}

func (this AttributeUnParse) readInfo(reader *ClassReader) {
	reader.readBytes(int(this.attrLen))
}
