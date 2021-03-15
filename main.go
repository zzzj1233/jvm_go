package main

import (
	"./classfile"
	"./classpath"
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
		if method.GetName() == "main" {
			Interpret(method)
		}

	}
}
