package classfile

type MemberInfo struct {
	AccessFlag uint16
	NameIndex  uint16
	DescIndex  uint16
	Attributes []AttributeInfo
	cp         ConstantPool
}

func (this *MemberInfo) GetName() string {
	return this.cp.getUtf8(this.NameIndex)
}

func (this *MemberInfo) GetDesc() string {
	return this.cp.getUtf8(this.DescIndex)
}

func newMemberInfo(reader *ClassReader, cp ConstantPool) *MemberInfo {
	info := &MemberInfo{}
	info.AccessFlag = reader.readUInt16()
	info.NameIndex = reader.readUInt16()
	info.DescIndex = reader.readUInt16()
	info.Attributes = readAttributes(reader, cp)
	info.cp = cp
	return info
}
