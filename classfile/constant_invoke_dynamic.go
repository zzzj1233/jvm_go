package classfile

type ConstantMethodHandle struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (this *ConstantMethodHandle) readInfo(reader *ClassReader, cp ConstantPool) {
	this.referenceKind = reader.readUInt8()
	this.referenceIndex = reader.readUInt16()
}

type ConstantMethodType struct {
	descIndex uint16
}

func (this *ConstantMethodType) readInfo(reader *ClassReader, cp ConstantPool) {
	this.descIndex = reader.readUInt16()
}

type ConstantInvokeDynamic struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (this *ConstantInvokeDynamic) readInfo(reader *ClassReader, cp ConstantPool) {
	this.bootstrapMethodAttrIndex = reader.readUInt16()
	this.nameAndTypeIndex = reader.readUInt16()
}
