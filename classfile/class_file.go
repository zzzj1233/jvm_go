package classfile

type ClassFile struct {
	Magic        uint32
	MajorVersion uint16
	SubVersion   uint16
	Pool         ConstantPool
	AccessFlag   uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []uint16
	Fields       []*MemberInfo
	Methods      []*MemberInfo
	Attributes   []AttributeInfo
}

func Parse(data []byte) *ClassFile {
	reader := &ClassReader{data: data}

	cf := &ClassFile{}

	cf.read(reader)

	return cf
}

func (this *ClassFile) read(reader *ClassReader) {
	this.readAndCheckMagic(reader)
	this.readAndCheckVersion(reader)
	this.Pool = readConstantPool(reader)
	this.readAccessFlag(reader)
	this.readThisClass(reader)
	this.readSuperClass(reader)
	this.readInterfaces(reader)
	this.readFields(reader)
	this.readMethods(reader)
	this.Attributes = readAttributes(reader, this.Pool)
}

func (this *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUInt32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: Magic!")
	}
	this.Magic = magic
}

func (this *ClassFile) readAndCheckVersion(reader *ClassReader) {
	this.SubVersion = reader.readUInt16()
	this.MajorVersion = reader.readUInt16()

	switch this.MajorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if this.SubVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (this *ClassFile) readAccessFlag(reader *ClassReader) {
	this.AccessFlag = reader.readUInt16()
}

func (this *ClassFile) readThisClass(reader *ClassReader) {
	this.ThisClass = reader.readUInt16()
}

func (this *ClassFile) readSuperClass(reader *ClassReader) {
	this.SuperClass = reader.readUInt16()
}

func (this *ClassFile) readInterfaces(reader *ClassReader) {
	this.Interfaces = reader.readUint16s()
}

func (this *ClassFile) readFields(reader *ClassReader) {
	fieldCount := reader.readUInt16()
	this.Fields = make([]*MemberInfo, 0, fieldCount)

	for i := 0; i < int(fieldCount); i++ {
		this.Fields = append(this.Fields, newMemberInfo(reader, this.Pool))
	}
}

func (this *ClassFile) readMethods(reader *ClassReader) {
	methodCount := reader.readUInt16()
	this.Methods = make([]*MemberInfo, 0, methodCount)

	for i := 0; i < int(methodCount); i++ {
		this.Methods = append(this.Methods, newMemberInfo(reader, this.Pool))
	}
}
