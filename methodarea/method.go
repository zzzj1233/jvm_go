package methodarea

type MethodFlag struct {
	*ClassMemberFlag
	synchronized bool
	native       bool
}

func (this *MethodFlag) IsSynchronized() bool {
	return this.synchronized
}

func (this *MethodFlag) IsNative() bool {
	return this.native
}

func newMethodFlag(flag int) *MethodFlag {
	return &MethodFlag{
		ClassMemberFlag: newClassMemberFlag(flag),
		synchronized:    flag&0x0020 != 0,
		native:          flag&0x0100 != 0,
	}
}

type Method struct {
	ClassMember
	flag      *MethodFlag
	maxStack  int
	maxLocals int
	code      []byte
}
