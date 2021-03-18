package methodarea

import (
	"../classfile"
)

type ClassMemberFlag struct {
	public    bool
	private   bool
	protected bool
	static    bool
	final     bool
}

func newClassMemberFlag(flag int) *ClassMemberFlag {
	return &ClassMemberFlag{
		public:    flag&0x0001 != 0,
		private:   flag&0x0002 != 0,
		protected: flag&0x0004 != 0,
		static:    flag&0x0008 != 0,
		final:     flag&0x0010 != 0,
	}
}

func (this *ClassMemberFlag) IsPublic() bool {
	return this.public
}

func (this *ClassMemberFlag) IsPrivate() bool {
	return this.private
}

func (this *ClassMemberFlag) IsProtected() bool {
	return this.protected
}

func (this *ClassMemberFlag) IsStatic() bool {
	return this.static
}

func (this *ClassMemberFlag) IsFinal() bool {
	return this.final
}

type FieldFlag struct {
	*ClassMemberFlag
	volatile  bool
	transient bool
	enum      bool
}

func newFieldFlag(flag int) *FieldFlag {
	return &FieldFlag{
		ClassMemberFlag: newClassMemberFlag(flag),
		volatile:        flag&0x0040 != 0,
		transient:       flag&0x0080 != 0,
		enum:            flag&0x4000 != 0,
	}
}

func (this *FieldFlag) IsVolatile() bool {
	return this.volatile
}

func (this *FieldFlag) IsTransient() bool {
	return this.transient
}

func (this *FieldFlag) IsEnum() bool {
	return this.enum
}

func newFields(cf *classfile.ClassFile, class *Class) []*Field {
	fields := make([]*Field, len(cf.Fields))

	for i, field := range cf.Fields {
		fields[i] = newField(field, class)
	}

	return fields
}

func newField(memberInfo *classfile.MemberInfo, class *Class) *Field {
	return &Field{
		ClassMember: ClassMember{
			Name:       memberInfo.GetName(),
			Descriptor: memberInfo.GetDesc(),
			Class:      class,
		},
		Flag:            newFieldFlag(int(memberInfo.AccessFlag)),
		FieldDescriptor: newDescriptor(memberInfo.GetDesc()),
	}
}

type ClassMember struct {
	Name       string
	Descriptor string
	Class      *Class
}

type Field struct {
	ClassMember
	Flag            *FieldFlag
	FieldDescriptor *Descriptor
}
