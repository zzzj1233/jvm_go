package methodarea

import (
	"../classfile"
)

type Class struct {
	flag            *ClassFlag
	className       string
	superClassName  string
	interfacesNames []string
	fields          []*Field
	methods         []*Method
	superClass      *Class
	interfaces      []*Class
	classLoader     *ClassLoader
}

func NewClass(cf *classfile.ClassFile, classLoader *ClassLoader) *Class {
	className := cf.GetClassName()
	superClassName := cf.GetSuperClassName()
	interfaceNames := cf.GetInterfaceNames()

	var superClass *Class = nil

	// 加载父类
	if className != "java/lang/Object" && superClassName != "" {
		superClass = classLoader.LoadClass(superClassName)
	}

	var interfaces = make([]*Class, len(cf.Interfaces))

	// 加载接口
	for idx, name := range interfaceNames {
		interfaces[idx] = classLoader.LoadClass(name)
	}

	clazz := &Class{
		flag:            newClassFlag(int(cf.AccessFlag)),
		className:       className,
		superClassName:  superClassName,
		interfacesNames: interfaceNames,
		fields:          nil,
		methods:         nil,
		superClass:      superClass,
		interfaces:      interfaces,
		classLoader:     classLoader,
	}

	clazz.methods = newMethods(cf, clazz)
	clazz.fields = newFields(cf, clazz)

	return clazz
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
