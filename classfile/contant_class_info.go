package classfile

type ConstantClassInfo struct {
	classNameIdx uint16
	pool         ConstantPool
}

func (this *ConstantClassInfo) readInfo(reader *ClassReader, cp ConstantPool) {
	this.classNameIdx = reader.readUInt16()
	this.pool = cp
}

func (this *ConstantClassInfo) getClassName() string {
	return this.pool.getUtf8(this.classNameIdx)
}
