package classfile

import (
	"strconv"
)

type ConstantInfo interface {
	readInfo(reader *ClassReader, cp ConstantPool)
}

func newConstantInfo(tag uint8) ConstantInfo {
	switch tag {

	case 1:
		return &ConstantUtf8Info{}
	case 3:
		return &ConstantIntegerInfo{}
	case 4:
		return &ConstantFloatInfo{}
	case 5:
		return &ConstantLongInfo{}
	case 6:
		return &ConstantDoubleInfo{}
	case 7:
		return &ConstantClassInfo{}
	case 8:
		return &ConstantStrInfo{}
	case 9:
		return &ConstantFieldRefInfo{}
	case 10:
		return &ConstantMethodRefInfo{}
	case 11:
		return &ConstantInterfaceMethodRefInfo{}
	case 12:
		return &ConstantNameType{}
	case 15:
		return &ConstantMethodHandle{}
	case 16:
		return &ConstantMethodType{}
	case 18:
		return &ConstantInvokeDynamic{}
	}

	panic("java.lang.ClassFormatError: illegal constant pool tag : " + strconv.Itoa(int(tag)))
}

func ReadConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUInt8()
	info := newConstantInfo(tag)
	info.readInfo(reader, cp)
	return info
}
