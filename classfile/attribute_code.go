package classfile

type CodeAttribute struct {
	cp             ConstantPool
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []*ExceptionTableEntry
	Attributes     []AttributeInfo
}

func (this *CodeAttribute) GetCode() []byte {
	return this.Code
}

func newCodeAttribute(cp ConstantPool) *CodeAttribute {
	return &CodeAttribute{cp: cp}
}

func (this *CodeAttribute) readInfo(reader *ClassReader) {
	this.MaxStack = reader.readUInt16()
	this.MaxLocals = reader.readUInt16()
	this.Code = reader.readBytes(int(reader.readUInt32()))
	exTableLen := reader.readUInt16()
	this.ExceptionTable = make([]*ExceptionTableEntry, 0, exTableLen)
	for i := 0; i < int(exTableLen); i++ {
		this.ExceptionTable = append(this.ExceptionTable, newExceptionTableEntry(reader))
	}
	this.Attributes = readAttributes(reader, this.cp)
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func newExceptionTableEntry(reader *ClassReader) *ExceptionTableEntry {
	entry := &ExceptionTableEntry{}
	entry.startPc = reader.readUInt16()
	entry.endPc = reader.readUInt16()
	entry.handlerPc = reader.readUInt16()
	entry.catchType = reader.readUInt16()
	return entry
}
