package methodarea

import (
	"../classfile"
	"../classpath"
	"strings"
)

type ClassLoader struct {
	classPool map[string]*Class
	cp        *classpath.ClassPath
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		classPool: make(map[string]*Class, 106),
		cp:        cp,
	}
}

func (this *ClassLoader) LoadClass(name string) *Class {
	class := this.classPool[name]

	if class != nil {
		return class
	}

	// 加载
	defineClass, cf := this.defineClass(name)

	this.classPool[name] = defineClass

	// 验证
	// verify(Class)

	// 初始化,执行static代码块,初始化静态变量
	prepare(defineClass, cf)

	return defineClass
}

func prepare(class *Class, cf *classfile.ClassFile) {
	staticAttributeMap := make(map[string]*Attribute, 4)
	instanceAttributeMap := make(map[string]*Field, 16)

	for _, field := range class.fields {

		if field.Flag.IsStatic() {
			staticAttributeMap[field.Name] = NewAttribute(field)
		} else {
			instanceAttributeMap[field.Name] = field
		}
	}

	class.StaticAttributes = staticAttributeMap
	class.InstanceFieldMap = instanceAttributeMap
}

func (this *ClassLoader) readClass(classFileName string) []byte {
	var bytes []byte
	var err error

	bytes, _, err = this.cp.BootClassPath.ReadClass(classFileName)

	if err == nil {
		return bytes
	}

	bytes, _, err = this.cp.ExtClassPath.ReadClass(classFileName)

	if err == nil {
		return bytes
	}

	bytes, _, err = this.cp.UsrClassPath.ReadClass(classFileName)

	if err == nil {
		return bytes
	}

	panic("read file error , can not find the file : " + classFileName)
}

func (this *ClassLoader) defineClass(className string) (*Class, *classfile.ClassFile) {
	classFileName := strings.ReplaceAll(className, ".", "/") + ".class"

	classData := this.readClass(classFileName)

	cf := classfile.Parse(classData)

	return NewClass(cf, this), cf
}
