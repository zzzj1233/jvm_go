package main

import (
	"./classpath"
	"./context"
	"./methodarea"
	"./rt"
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

func startJvm(c *Cmd) {
	cp := classpath.Parse(c.jreOptions, c.cpOptions)

	context.Loader = methodarea.NewClassLoader(cp)

	class := context.Loader.LoadClass(c.classname)

	thread := rt.NewThread("main")

	method := class.GetMainMethod()

	Interpret(method, thread)
}
