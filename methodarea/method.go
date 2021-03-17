package methodarea

import (
	"../classfile"
)

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

func newMethods(cf *classfile.ClassFile, class *Class) []*Method {
	methods := make([]*Method, len(cf.Methods))

	for i, method := range cf.Methods {
		methods[i] = newMethod(method, class)
	}

	return methods
}

func newMethod(memberInfo *classfile.MemberInfo, class *Class) *Method {

	var maxStack uint16
	var maxLocals uint16
	var code []byte

	for _, attr := range memberInfo.Attributes {

		codeAttr, ok := attr.(*classfile.CodeAttribute)

		if ok {
			maxStack = codeAttr.MaxStack
			maxLocals = codeAttr.MaxLocals
			code = codeAttr.Code
		}

	}

	return &Method{
		ClassMember: ClassMember{
			Name:       memberInfo.GetName(),
			Descriptor: memberInfo.GetDesc(),
			Class:      class,
		},
		flag:      newMethodFlag(int(memberInfo.AccessFlag)),
		maxStack:  maxStack,
		maxLocals: maxLocals,
		code:      code,
	}
}

type Method struct {
	ClassMember
	flag      *MethodFlag
	maxStack  uint16
	maxLocals uint16
	code      []byte
}
