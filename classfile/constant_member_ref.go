package classfile

type ConstantFieldRefInfo struct {
	ConstantMemberInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantMemberInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
	pool             ConstantPool
}

func (this *ConstantMemberInfo) readInfo(reader *ClassReader, cp ConstantPool) {
	this.classIndex = reader.readUInt16()
	this.nameAndTypeIndex = reader.readUInt16()
	this.pool = cp
}

func (this *ConstantMemberInfo) GetClassName() string {
	return this.pool.getClassInfo(this.classIndex).getClassName()
}

func (this *ConstantMemberInfo) GetNameAndType() (string, string) {
	info := this.pool.getNameAndTypeInfo(this.nameAndTypeIndex)
	return info.GetName(), info.GetDesc()
}
