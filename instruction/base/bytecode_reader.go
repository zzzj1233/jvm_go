package base

type BytecodeReader struct {
	pc   uint
	code []byte
}

func NewBytecodeReader(pc uint, code []byte) *BytecodeReader {
	return &BytecodeReader{
		pc:   pc,
		code: code,
	}
}

func (this *BytecodeReader) ReadUInt8() uint8 {
	operands := this.code[this.pc]
	this.pc++
	return operands
}

func (this *BytecodeReader) ReadInt8() int8 {
	return int8(this.ReadUInt8())
}

func (this *BytecodeReader) ReadUInt16() uint16 {
	return uint16(this.ReadUInt8())<<8 | uint16(this.ReadUInt8())
}

func (this *BytecodeReader) ReadInt16() int16 {
	return int16(int16(this.ReadInt8()<<8) | int16(this.ReadInt8()))
}

func (this *BytecodeReader) ReadUInt32() uint32 {
	byte1 := uint32(this.ReadUInt8())
	byte2 := uint32(this.ReadUInt8())
	byte3 := uint32(this.ReadUInt8())
	byte4 := uint32(this.ReadUInt8())

	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (this *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(this.ReadUInt8())
	byte2 := int32(this.ReadUInt8())
	byte3 := int32(this.ReadUInt8())
	byte4 := int32(this.ReadUInt8())

	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (this *BytecodeReader) SetNextPc(pc uint) {
	this.pc += pc
}
