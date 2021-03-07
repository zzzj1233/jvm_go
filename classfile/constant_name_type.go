package classfile

type ConstantNameType struct {
	nameIndex uint16
	descIndex uint16
	pool      ConstantPool
}

func (this *ConstantNameType) readInfo(reader *ClassReader, cp ConstantPool) {
	this.nameIndex = reader.readUInt16()
	this.descIndex = reader.readUInt16()
	this.pool = cp
}

func (this *ConstantNameType) GetName() string {
	return this.pool.getUtf8(this.nameIndex)
}

func (this *ConstantNameType) GetDesc() string {
	return this.pool.getUtf8(this.descIndex)
}
