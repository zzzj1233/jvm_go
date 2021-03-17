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
	class = this.defineClass(name)

	this.classPool[name] = class

	// 验证
	// verify(Class)

	// 初始化

	return class
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

func (this *ClassLoader) defineClass(className string) *Class {
	classFileName := strings.ReplaceAll(className, ".", "/") + ".Class"

	classData := this.readClass(classFileName)

	cf := classfile.Parse(classData)

	return NewClass(cf, this)
}
