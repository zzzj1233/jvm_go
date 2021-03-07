package classfile

type ConstantStrInfo struct {
	utf8Idx uint16
	pool    ConstantPool
}

func (this *ConstantStrInfo) readInfo(reader *ClassReader, cp ConstantPool) {
	this.utf8Idx = reader.readUInt16()
	this.pool = cp
}

func (this *ConstantStrInfo) GetStr() string {
	return this.pool.getUtf8(this.utf8Idx)
}
