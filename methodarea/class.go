package methodarea

import (
	"../classfile"
	"strings"
)

type Class struct {
	flag             *ClassFlag
	className        string
	superClassName   string
	packageName      string
	interfacesNames  []string
	fields           []*Field
	methods          []*Method
	superClass       *Class
	interfaces       []*Class
	classLoader      *ClassLoader
	StaticAttributes map[string]*Attribute
	InstanceFieldMap map[string]*Field
	Pool             classfile.ConstantPool
}

func (this *Class) GetFields() []*Field {
	return this.fields
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

	var packageName string

	lastIndex := strings.LastIndex(className, "/") + 1

	if lastIndex != 0 {
		packageName = strings.ReplaceAll(className[:lastIndex-1], "/", ".")
	}

	clazz := &Class{
		flag:            newClassFlag(int(cf.AccessFlag)),
		className:       className,
		superClassName:  superClassName,
		interfacesNames: interfaceNames,
		packageName:     packageName,
		fields:          nil,
		methods:         nil,
		superClass:      superClass,
		interfaces:      interfaces,
		classLoader:     classLoader,
		Pool:            cf.Pool,
	}

	clazz.methods = newMethods(cf, clazz)
	clazz.fields = newFields(cf, clazz)

	return clazz
}

func (this *Class) IsAssignableFrom(other *Class) bool {
	if this == other {
		return true
	}

	// 1. 如果当前class是类
	if !this.flag.IsInterface() {

		if other.flag.IsInterface() {
			return false
		}

		var super = other.superClass

		for super != nil {
			if super == this {
				return true
			}
			super = super.superClass
		}

		// 2. 如果当前class是interface
	} else {

		// other是interface
		if other.flag.IsInterface() {
			var super = other.superClass

			for super != nil {
				if super == this {
					return true
				}
				super = super.superClass
			}
			// other是class
		} else {
			// 2.1 other继承的任何类是否实现了接口
			var super = other.superClass

			for super != nil {
				if this.IsAssignableFrom(super) {
					return true
				}
				super = super.superClass
			}
			// 2.2 other实现的任何类是否实现了接口
			interfaces := other.interfaces

			for _, interFace := range interfaces {
				if this.IsAssignableFrom(interFace) {
					return true
				}
			}
		}
	}
	// 3. 当前class是数组
	return false
}

func (this *Class) AccessStaticField(visitor *Class, name string) *Attribute {
	attribute := this.StaticAttributes[name]

	if attribute == nil {
		panic("java.lang.NoSuchFieldError")
	}

	this.Access(visitor, attribute.Field.Flag.ClassMemberFlag)

	return attribute
}

func (this *Class) Access(visitor *Class, flag *ClassMemberFlag) {
	if flag.IsProtected() {
		if this != visitor {
			// 递归找
			temp := visitor
			for temp.GetSuperClassName() != this.GetClassName() {
				if temp.GetSuperClassName() == "" {
					panic("java.lang.IllegalAccessError")
				}
				temp = temp.GetSuperClass()
			}
		}
	} else if flag.IsPrivate() && this != visitor {
		panic("java.lang.IllegalAccessError")
	} else if flag.IsPublic() {
		// no thing
	} else { // default package
		if this.GetPackageName() != visitor.GetPackageName() {
			panic("java.lang.IllegalAccessError")
		}
	}
}

func (this *Class) GetSuperClass() *Class {
	return this.superClass
}

func (this *Class) GetPackageName() string {
	return this.packageName
}

func (this *Class) GetSuperClassName() string {
	return this.superClassName
}

func (this *Class) GetClassName() string {
	return this.className
}

func (this *Class) GetMainMethod() *Method {
	for _, method := range this.methods {
		if method.Name == "main" {
			return method
		}
	}
	panic("java.lang.NoSuchMethodError")
}

type ClassFlag struct {
	public     bool
	final      bool
	abstract   bool
	annotation bool
	enum       bool
	interFace  bool
}

func newClassFlag(flag int) *ClassFlag {
	return &ClassFlag{
		public:     flag&0x0001 != 0,
		final:      flag&0x0010 != 0,
		abstract:   flag&0x0040 != 0,
		annotation: flag&0x2000 != 0,
		enum:       flag&0x4000 != 0,
		interFace:  flag&0x0200 != 0,
	}
}

func (this *ClassFlag) IsInterface() bool {
	return this.interFace
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
