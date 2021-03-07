package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := reader.readUInt16()

	attrs := make([]AttributeInfo, 0, attributeCount)

	for i := 0; i < int(attributeCount); i++ {
		attrs = append(attrs, readAttribute(reader, cp))
	}

	return attrs
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrIdx := reader.readUInt16()
	attributeName := cp.getUtf8(attrIdx)
	attributeLen := reader.readUInt32()
	attribute := newAttribute(attributeName, attributeLen, cp)
	attribute.readInfo(reader)
	return attribute
}

func newAttribute(attributeName string, attributeLen uint32, cp ConstantPool) AttributeInfo {
	switch attributeName {
	case "Code":
		return newCodeAttribute(cp)
	case "Exceptions":
		return newAttributeExceptions(cp)
	case "Deprecated":
		return &AttributeDeprecated{}
	case "Synthetic":
		return &AttributeSynthetic{}
	case "AttributeSourceFile":
		return &AttributeSourceFile{cp: cp}
	case "LocalVariableTable":
		return &AttributeLocalVarTable{}
	case "LineNumberTable":
		return &AttributeLineNumberTable{}
	}

	return &AttributeUnParse{attrLen: attributeLen}
}
