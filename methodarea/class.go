package methodarea

import (
	"../classfile"
	"fmt"
)

type Class struct {
	flag            *ClassFlag
	className       string
	superClassName  string
	interfacesNames string
	fields          []Field
	methods         []Method
	superClass      *Class
	interfaces      []*Class
	classLoader     *ClassLoader
}

func NewClass(cf *classfile.ClassFile) *Class {
	fmt.Println(cf.GetClassName())
	return &Class{
		flag:            newClassFlag(int(cf.AccessFlag)),
		className:       cf.GetClassName(),
		superClassName:  "",
		interfacesNames: "",
		fields:          nil,
		methods:         nil,
		superClass:      nil,
		interfaces:      nil,
		classLoader:     nil,
	}
}

type ClassFlag struct {
	public     bool
	final      bool
	abstract   bool
	annotation bool
	enum       bool
}

func newClassFlag(flag int) *ClassFlag {
	return &ClassFlag{
		public:     flag&0x0001 != 0,
		final:      flag&0x0010 != 0,
		abstract:   flag&0x0040 != 0,
		annotation: flag&0x2000 != 0,
		enum:       flag&0x4000 != 0,
	}
}

func (this *ClassFlag) IsPublic() bool {
	return this.public
}

func (this *ClassFlag) IsFinal() bool {
	return this.final
}

func (this *ClassFlag) IsAbstract() bool {
	return this.abstract
}

func (this *ClassFlag) IsAnnotation() bool {
	return this.annotation
}

func (this *ClassFlag) IsEnum() bool {
	return this.enum
}
