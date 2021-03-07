package main

import (
	"./classfile"
	"./classpath"
	"./instruction"
	"./instruction/base"
	"./rt"
	"strings"
)

func main() {

	c := parseCommand()

	if c.printVersion {
		c.version()
	} else if c.printHelp {
		c.help()
	} else {
		startJvm(c)
	}

}

func getFileBytes(classFileName string, cp *classpath.ClassPath) []byte {
	var bytes []byte
	var err error

	bytes, _, err = cp.BootClassPath.ReadClass(classFileName)

	if err == nil {
		return bytes
	}

	bytes, _, err = cp.ExtClassPath.ReadClass(classFileName)

	if err == nil {
		return bytes
	}

	bytes, _, err = cp.UsrClassPath.ReadClass(classFileName)

	if err == nil {
		return bytes
	}

	panic("read file error , can not find the file : " + classFileName)
}

func startJvm(c *Cmd) {
	cp := classpath.Parse(c.jreOptions, c.cpOptions)

	classFileName := strings.ReplaceAll(c.classname, ".", "/") + ".class"

	bytes := getFileBytes(classFileName, cp)

	cf := classfile.Parse(bytes)

	for _, method := range cf.Methods {
		for _, attr := range method.Attributes {
			codeAttribute := attr.(*classfile.CodeAttribute)

			if method.GetName() == "test" {
				reader := base.NewBytecodeReader(0, codeAttribute.GetCode())
				frame := rt.NewStackFrame(rt.NewLocalVarTable(codeAttribute.MaxLocals), rt.NewOperateStack(codeAttribute.MaxStack))
				instruction.Loop(reader, frame)
			}
		}

	}
}
