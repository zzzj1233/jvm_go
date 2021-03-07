package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (this *ClassReader) readUInt8() uint8 {
	value := this.data[0]
	this.data = this.data[1:]
	return value
}

func (this *ClassReader) readUInt16() uint16 {
	value := binary.BigEndian.Uint16(this.data)
	this.data = this.data[2:]
	return value
}

func (this *ClassReader) readUInt32() uint32 {
	value := binary.BigEndian.Uint32(this.data)
	this.data = this.data[4:]
	return value
}

func (this *ClassReader) readUInt64() uint64 {
	value := binary.BigEndian.Uint64(this.data)
	this.data = this.data[8:]
	return value
}

func (this *ClassReader) readUint16s() []uint16 {
	n := this.readUInt16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = this.readUInt16()
	}
	return s
}

func (this *ClassReader) readBytes(n int) []byte {
	val := this.data[:n]
	this.data = this.data[n:]
	return val
}
