package classfile

type AttributeExceptions struct {
	cp                  ConstantPool
	exceptionIndexTable []uint16
}

func newAttributeExceptions(cp ConstantPool) *AttributeExceptions {
	return &AttributeExceptions{cp: cp}
}

func (this AttributeExceptions) readInfo(reader *ClassReader) {
	this.exceptionIndexTable = reader.readUint16s()
}
