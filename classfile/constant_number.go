package classfile

import "math"

type ConstantIntegerInfo struct {
	Value int32
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader, cp ConstantPool) {
	this.Value = int32(reader.readUInt32())
}

type ConstantLongInfo struct {
	Value int64
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader, cp ConstantPool) {
	this.Value = int64(reader.readUInt64())
}

type ConstantFloatInfo struct {
	Value float32
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader, cp ConstantPool) {
	bytes := reader.readUInt32()
	this.Value = math.Float32frombits(bytes)
}

type ConstantDoubleInfo struct {
	Value float64
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader, cp ConstantPool) {
	bytes := reader.readUInt64()
	this.Value = math.Float64frombits(bytes)
}
