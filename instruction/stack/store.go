package stack

import "../base"
import "../../rt"

// AStore
type AStore struct {
	base.Index8Instruction
}

func _AStoreExecute(index int, frame *rt.StackFrame) {
	frame.LocalVarTable.PutRef(index, frame.OperateStack.PopObj())
}

func (this AStore) Execute(frame *rt.StackFrame) {
	_AStoreExecute(this.Index, frame)
}

type AStore0 struct {
	base.NoOperandsInstruction
}

func (this AStore0) Execute(frame *rt.StackFrame) {
	_AStoreExecute(0, frame)
}

type AStore1 struct {
	base.NoOperandsInstruction
}

func (this AStore1) Execute(frame *rt.StackFrame) {
	_AStoreExecute(1, frame)
}

type AStore2 struct {
	base.NoOperandsInstruction
}

func (this AStore2) Execute(frame *rt.StackFrame) {
	_AStoreExecute(2, frame)
}

type AStore3 struct {
	base.NoOperandsInstruction
}

func (this AStore3) Execute(frame *rt.StackFrame) {
	_AStoreExecute(3, frame)
}

// IStore
type IStore struct {
	base.Index8Instruction
}

func _IStoreExecute(index int, frame *rt.StackFrame) {
	frame.LocalVarTable.PutInt(index, frame.OperateStack.PopInt())
}

func (this IStore) Execute(frame *rt.StackFrame) {
	_IStoreExecute(this.Index, frame)
}

type IStore0 struct {
	base.NoOperandsInstruction
}

func (this IStore0) Execute(frame *rt.StackFrame) {
	_IStoreExecute(0, frame)
}

type IStore1 struct {
	base.NoOperandsInstruction
}

func (this IStore1) Execute(frame *rt.StackFrame) {
	_IStoreExecute(1, frame)
}

type IStore2 struct {
	base.NoOperandsInstruction
}

func (this IStore2) Execute(frame *rt.StackFrame) {
	_IStoreExecute(2, frame)
}

type IStore3 struct {
	base.NoOperandsInstruction
}

func (this IStore3) Execute(frame *rt.StackFrame) {
	_IStoreExecute(3, frame)
}

// LStore
type LStore struct {
	base.Index8Instruction
}

func _LStoreExecute(index int, frame *rt.StackFrame) {
	frame.LocalVarTable.PutLong(index, frame.OperateStack.PopLong())
}

func (this LStore) Execute(frame *rt.StackFrame) {
	_LStoreExecute(this.Index, frame)
}

type LStore0 struct {
	base.NoOperandsInstruction
}

func (this LStore0) Execute(frame *rt.StackFrame) {
	_LStoreExecute(0, frame)
}

type LStore1 struct {
	base.NoOperandsInstruction
}

func (this LStore1) Execute(frame *rt.StackFrame) {
	_LStoreExecute(1, frame)
}

type LStore2 struct {
	base.NoOperandsInstruction
}

func (this LStore2) Execute(frame *rt.StackFrame) {
	_LStoreExecute(2, frame)
}

type LStore3 struct {
	base.NoOperandsInstruction
}

func (this LStore3) Execute(frame *rt.StackFrame) {
	_LStoreExecute(3, frame)
}

// DStore
type DStore struct {
	base.Index8Instruction
}

func _DStoreExecute(index int, frame *rt.StackFrame) {
	frame.LocalVarTable.PutDouble(index, frame.OperateStack.PopDouble())
}

func (this DStore) Execute(frame *rt.StackFrame) {
	_DStoreExecute(this.Index, frame)
}

type DStore0 struct {
	base.NoOperandsInstruction
}

func (this DStore0) Execute(frame *rt.StackFrame) {
	_DStoreExecute(0, frame)
}

type DStore1 struct {
	base.NoOperandsInstruction
}

func (this DStore1) Execute(frame *rt.StackFrame) {
	_DStoreExecute(1, frame)
}

type DStore2 struct {
	base.NoOperandsInstruction
}

func (this DStore2) Execute(frame *rt.StackFrame) {
	_DStoreExecute(2, frame)
}

type DStore3 struct {
	base.NoOperandsInstruction
}

func (this DStore3) Execute(frame *rt.StackFrame) {
	_DStoreExecute(3, frame)
}

// FStore
type FStore struct {
	base.Index8Instruction
}

func _FStoreExecute(index int, frame *rt.StackFrame) {
	frame.LocalVarTable.PutFloat(index, frame.OperateStack.PopFloat())
}

func (this FStore) Execute(frame *rt.StackFrame) {
	_FStoreExecute(this.Index, frame)
}

type FStore0 struct {
	base.NoOperandsInstruction
}

func (this FStore0) Execute(frame *rt.StackFrame) {
	_FStoreExecute(0, frame)
}

type FStore1 struct {
	base.NoOperandsInstruction
}

func (this FStore1) Execute(frame *rt.StackFrame) {
	_FStoreExecute(1, frame)
}

type FStore2 struct {
	base.NoOperandsInstruction
}

func (this FStore2) Execute(frame *rt.StackFrame) {
	_FStoreExecute(2, frame)
}

type FStore3 struct {
	base.NoOperandsInstruction
}

func (this FStore3) Execute(frame *rt.StackFrame) {
	_FStoreExecute(3, frame)
}
