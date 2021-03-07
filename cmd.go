package main

import (
	"flag"
	"os"
)

const version = "1.0"

type Cmd struct {
	printHelp    bool
	printVersion bool
	args         []string
	jreOptions   string
	cpOptions    string
	classname    string
}

func parseCommand() *Cmd {
	c := &Cmd{}

	flag.BoolVar(&c.printHelp, "?", false, "查看帮助")
	flag.BoolVar(&c.printHelp, "help", false, "查看帮助")
	flag.BoolVar(&c.printVersion, "v", false, "查看版本")
	flag.BoolVar(&c.printVersion, "V", false, "查看版本")
	flag.StringVar(&c.jreOptions, "jre", "", "指定jre路径")
	flag.StringVar(&c.cpOptions, "cp", "", "指定classpath路径")
	flag.StringVar(&c.cpOptions, "classpath", "", "指定classpath路径")
	flag.Parse()

	if len(os.Args) == 1 {
		panic("must special class file")
	}

	c.classname = os.Args[1]
	c.args = os.Args[2:]

	return c
}

func (this Cmd) help() {
	println("-? | -help 查看帮助")
	println("-v | -V 查看版本")
	println("-jre 指定jre路径")
	println("-cp | -classpath 指定classpath路径")
}

func (this Cmd) version() {
	println("jvm by zzzj , version = ", version)
}
