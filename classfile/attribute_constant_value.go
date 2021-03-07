package classfile

type AttributeConstantValue struct {
	value string
	cp    ConstantPool
}

func (this AttributeConstantValue) readInfo(reader *ClassReader) {
	this.value = this.cp.getUtf8(reader.readUInt16())
}
