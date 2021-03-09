package math

import "../base"
import "../../rt"

// 给局部变量表指定index变量incr指定const的值
type IInc struct {
	index    uint8
	constInc int8
}

func (this *IInc) FetchOperands(reader *base.BytecodeReader) {
	this.index = reader.ReadUInt8()
	this.constInc = reader.ReadInt8()
}

func (this *IInc) Execute(frame *rt.StackFrame) {
	val := frame.LocalVarTable.GetInt(int(this.index))
	frame.LocalVarTable.PutInt(int(this.index), val+int32(this.constInc))
}
