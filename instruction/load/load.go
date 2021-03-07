package load

import "../base"
import "../../rt"

// ALOAD
type ALoad struct {
	_ALoad
}

type _ALoad struct {
	index int
}

func (this *_ALoad) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushObj(frame.LocalVarTable.GetRef(this.index))
}

func (this *ALoad) FetchOperands(reader *base.BytecodeReader) {
	this._ALoad.index = int(reader.ReadInt8())
}

type ALoad0 struct {
	_ALoad
	base.NoOperandsInstruction
}

type ALoad1 struct {
	_ALoad
}

func (this *ALoad1) FetchOperands(reader *base.BytecodeReader) {
	this._ALoad.index = 1
}

type ALoad2 struct {
	_ALoad
}

func (this *ALoad2) FetchOperands(reader *base.BytecodeReader) {
	this._ALoad.index = 2
}

type ALoad3 struct {
	_ALoad
}

func (this *ALoad3) FetchOperands(reader *base.BytecodeReader) {
	this._ALoad.index = 3
}

// DLOAD
type DLoad struct {
	_DLoad
}

type _DLoad struct {
	index int
}

func (this *_DLoad) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushDouble(frame.LocalVarTable.GetDouble(this.index))
}

func (this *DLoad) FetchOperands(reader *base.BytecodeReader) {
	this._DLoad.index = int(reader.ReadInt8())
}

type DLoad0 struct {
	_DLoad
	base.NoOperandsInstruction
}

type DLoad1 struct {
	_DLoad
}

func (this *DLoad1) FetchOperands(reader *base.BytecodeReader) {
	this._DLoad.index = 1
}

type DLoad2 struct {
	_DLoad
}

func (this *DLoad2) FetchOperands(reader *base.BytecodeReader) {
	this._DLoad.index = 2
}

type DLoad3 struct {
	_DLoad
}

func (this *DLoad3) FetchOperands(reader *base.BytecodeReader) {
	this._DLoad.index = 3
}

// LLoad

type LLoad struct {
	_LLoad
}

type _LLoad struct {
	index int
}

func (this *_LLoad) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushLong(frame.LocalVarTable.GetLong(this.index))
}

func (this *LLoad) FetchOperands(reader *base.BytecodeReader) {
	this._LLoad.index = int(reader.ReadInt8())
}

type LLoad0 struct {
	_LLoad
	base.NoOperandsInstruction
}

type LLoad1 struct {
	_LLoad
}

func (this *LLoad1) FetchOperands(reader *base.BytecodeReader) {
	this._LLoad.index = 1
}

type LLoad2 struct {
	_LLoad
}

func (this *LLoad2) FetchOperands(reader *base.BytecodeReader) {
	this._LLoad.index = 2
}

type LLoad3 struct {
	_LLoad
}

func (this *LLoad3) FetchOperands(reader *base.BytecodeReader) {
	this._LLoad.index = 3
}

// ILOAD
type ILoad struct {
	_ILoad
}

type _ILoad struct {
	index int
}

func (this *_ILoad) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushInt(frame.LocalVarTable.GetInt(this.index))
}

func (this *ILoad) FetchOperands(reader *base.BytecodeReader) {
	this._ILoad.index = int(reader.ReadInt8())
}

type ILoad0 struct {
	_ILoad
	base.NoOperandsInstruction
}

type ILoad1 struct {
	_ILoad
}

func (this *ILoad1) FetchOperands(reader *base.BytecodeReader) {
	this._ILoad.index = 1
}

type ILoad2 struct {
	_ILoad
}

func (this *ILoad2) FetchOperands(reader *base.BytecodeReader) {
	this._ILoad.index = 2
}

type ILoad3 struct {
	_ILoad
}

func (this *ILoad3) FetchOperands(reader *base.BytecodeReader) {
	this._ILoad.index = 3
}

// FLOAD
type FLoad struct {
	_FLoad
}

type _FLoad struct {
	index int
}

func (this *_FLoad) Execute(frame *rt.StackFrame) {
	frame.OperateStack.PushFloat(frame.LocalVarTable.GetFloat(this.index))
}

func (this *FLoad) FetchOperands(reader *base.BytecodeReader) {
	this._FLoad.index = int(reader.ReadInt8())
}

type FLoad0 struct {
	_FLoad
	base.NoOperandsInstruction
}

type FLoad1 struct {
	_FLoad
}

func (this *FLoad1) FetchOperands(reader *base.BytecodeReader) {
	this._FLoad.index = 1
}

type FLoad2 struct {
	_FLoad
}

func (this *FLoad2) FetchOperands(reader *base.BytecodeReader) {
	this._FLoad.index = 2
}

type FLoad3 struct {
	_FLoad
}

func (this *FLoad3) FetchOperands(reader *base.BytecodeReader) {
	this._FLoad.index = 3
}
