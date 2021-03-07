package instruction

import (
	"../rt"
)
import "./base"
import "./constant"
import "./push"
import "./store"
import "./load"

func NewInstruction(instructionCode uint8) base.Instruction {
	switch instructionCode {
	// constant
	case 0x1:
		return &constant.AConstNull{}
	case 0xe:
		return &constant.DConst0{}
	case 0xf:
		return &constant.DConst1{}
	case 0xb:
		return &constant.FConst0{}
	case 0xc:
		return &constant.FConst1{}
	case 0xd:
		return &constant.FConst2{}
	case 0x2:
		return &constant.IConstM1{}
	case 0x3:
		return &constant.IConst0{}
	case 0x4:
		return &constant.IConst1{}
	case 0x5:
		return &constant.IConst2{}
	case 0x6:
		return &constant.IConst3{}
	case 0x7:
		return &constant.IConst4{}
	case 0x8:
		return &constant.IConst5{}
	case 0x9:
		return &constant.LConst0{}
	case 0xa:
		return &constant.LConst1{}
	// push
	case 0x10:
		return &push.BiPush{}
	case 0x11:
		return &push.SiPush{}
	// load
	case 0x19:
		return &load.ALoad{}
	case 0x2a:
		return &load.ALoad0{}
	case 0x2b:
		return &load.ALoad1{}
	case 0x2c:
		return &load.ALoad2{}
	case 0x2d:
		return &load.ALoad3{}
	case 0x18:
		return &load.DLoad{}
	case 0x26:
		return &load.DLoad0{}
	case 0x27:
		return &load.DLoad1{}
	case 0x28:
		return &load.DLoad2{}
	case 0x29:
		return &load.DLoad3{}
	case 0x17:
		return &load.FLoad{}
	case 0x22:
		return &load.FLoad0{}
	case 0x23:
		return &load.FLoad1{}
	case 0x24:
		return &load.FLoad2{}
	case 0x25:
		return &load.FLoad3{}
	case 0x15:
		return &load.ILoad{}
	case 0x1a:
		return &load.ILoad0{}
	case 0x1b:
		return &load.ILoad1{}
	case 0x1c:
		return &load.ILoad2{}
	case 0x1d:
		return &load.ILoad3{}
	case 0x16:
		return &load.LLoad{}
	case 0x1e:
		return &load.LLoad0{}
	case 0x1f:
		return &load.LLoad1{}
	case 0x20:
		return &load.LLoad2{}
	case 0x21:
		return &load.LLoad3{}
	// store
	case 0x3a:
		return &store.AStore{}
	case 0x4b:
		return &store.AStore0{}
	case 0x4c:
		return &store.AStore1{}
	case 0x4d:
		return &store.AStore2{}
	case 0x4e:
		return &store.AStore3{}
	case 0x39:
		return &store.DStore{}
	case 0x47:
		return &store.DStore0{}
	case 0x48:
		return &store.DStore1{}
	case 0x49:
		return &store.DStore2{}
	case 0x4a:
		return &store.DStore3{}
	case 0x38:
		return &store.FStore{}
	case 0x43:
		return &store.FStore0{}
	case 0x44:
		return &store.FStore1{}
	case 0x45:
		return &store.FStore2{}
	case 0x46:
		return &store.FStore3{}
	case 0x36:
		return &store.IStore{}
	case 0x3b:
		return &store.IStore0{}
	case 0x3c:
		return &store.IStore1{}
	case 0x3d:
		return &store.IStore2{}
	case 0x3e:
		return &store.IStore3{}
	case 0x37:
		return &store.LStore{}
	case 0x3f:
		return &store.LStore0{}
	case 0x40:
		return &store.LStore1{}
	case 0x41:
		return &store.LStore2{}
	case 0x42:
		return &store.LStore3{}
	case 0xb1:
		return &base.NopInstruction{}
	}

	return nil
}

func Loop(reader *base.BytecodeReader, frame *rt.StackFrame) {
	var instr base.Instruction

	for {
		instructionCode := reader.ReadUInt8()

		instr = NewInstruction(instructionCode)

		_, ok := instr.(*base.NopInstruction)

		if ok {
			return
		}

		instr.FetchOperands(reader)

		instr.Execute(frame)
	}

}
