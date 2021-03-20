package main

import (
	"./classpath"
	"./context"
	"./methodarea"
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

	method := class.GetMainMethod()

	Interpret(method)
}
