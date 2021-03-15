package instruction

import "./base"
import "./constant"
import "./push"
import "./store"
import "./load"
import "./cast"
import "./condition"
import "./dup"
import "./math"
import "./pop"

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
	case 0x90:
		return &cast.D2F{}
	case 0x8e:
		return &cast.D2I{}
	case 0x8f:
		return &cast.D2L{}
	case 0x8d:
		return &cast.F2D{}
	case 0x8b:
		return &cast.F2I{}
	case 0x8c:
		return &cast.F2L{}
	case 0x91:
		return &cast.I2B{}
	case 0x92:
		return &cast.I2C{}
	case 0x87:
		return &cast.I2D{}
	case 0x86:
		return &cast.I2F{}
	case 0x85:
		return &cast.I2L{}
	case 0x93:
		return &cast.I2S{}
	case 0x8a:
		return &cast.L2D{}
	case 0x89:
		return &cast.L2F{}
	case 0x88:
		return &cast.L2I{}
	case 0xa5:
		return &condition.AcmpEq{}
	case 0xa6:
		return &condition.AcmpNe{}
	case 0xc7:
		return &condition.IfNonNull{}
	case 0xc6:
		return &condition.IfNull{}
	case 0x98:
		return &condition.DCmpG{}
	case 0x97:
		return &condition.DCmpL{}
	case 0x96:
		return &condition.FCmpG{}
	case 0x95:
		return &condition.FCmpL{}
	case 0x94:
		return &condition.LCmp{}
	case 0x9f:
		return &condition.IcmpEq{}
	case 0xa0:
		return &condition.IcmpNe{}
	case 0xa1:
		return &condition.IcmpLt{}
	case 0xa2:
		return &condition.IcmpGe{}
	case 0xa3:
		return &condition.IcmpGt{}
	case 0xa4:
		return &condition.IcmpLe{}
	case 0x99:
		return &condition.IfEq{}
	case 0x9a:
		return &condition.IfNe{}
	case 0x9b:
		return &condition.IfLt{}
	case 0x9c:
		return &condition.IfGe{}
	case 0x9d:
		return &condition.IfGt{}
	case 0x9e:
		return &condition.IfLe{}
	case 0xa7:
		return &condition.Goto{}
	case 0x59:
		return &dup.Dup{}
	case 0x5a:
		return &dup.DupX1{}
	case 0x5b:
		return &dup.DupX2{}
	case 0x5c:
		return &dup.Dup2{}
	case 0x5d:
		return &dup.Dup2X1{}
	case 0x5e:
		return &dup.Dup2X2{}
	case 0x63:
		return &math.DAdd{}
	case 0x60:
		return &math.IAdd{}
	case 0x62:
		return &math.FAdd{}
	case 0x61:
		return &math.LAdd{}
	case 0x7e:
		return &math.IAnd{}
	case 0x7f:
		return &math.LAnd{}
	case 0x6f:
		return &math.DDiv{}
	case 0x6e:
		return &math.FDiv{}
	case 0x6c:
		return &math.IDiv{}
	case 0x6d:
		return &math.LDiv{}
	case 0x84:
		return &math.IInc{}
	case 0x6b:
		return &math.DMul{}
	case 0x6a:
		return &math.FMul{}
	case 0x68:
		return &math.IMul{}
	case 0x69:
		return &math.LMul{}
	case 0x77:
		return &math.DNeg{}
	case 0x76:
		return &math.FNeg{}
	case 0x74:
		return &math.INeg{}
	case 0x75:
		return &math.LNeg{}
	case 0x80:
		return &math.IOr{}
	case 0x81:
		return &math.LOr{}
	case 0x73:
		return &math.DRem{}
	case 0x72:
		return &math.FRem{}
	case 0x70:
		return &math.IRem{}
	case 0x71:
		return &math.LRem{}
	case 0x78:
		return &math.IShl{}
	case 0x7a:
		return &math.IShR{}
	case 0x7c:
		return &math.IUShR{}
	case 0x79:
		return &math.LShl{}
	case 0x7b:
		return &math.LShr{}
	case 0x7d:
		return &math.LUShr{}
	case 0x67:
		return &math.DSub{}
	case 0x66:
		return &math.FSub{}
	case 0x64:
		return &math.ISub{}
	case 0x65:
		return &math.LSub{}
	case 0x82:
		return &math.IXor{}
	case 0x83:
		return &math.LXor{}
	case 0x57:
		return &pop.Pop{}
	case 0x58:
		return &pop.Pop2{}
	}
	return &base.NopInstruction{}
}
