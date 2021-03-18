package main

import (
	"./classpath"
	"./methodarea"
	"./rt"
	"fmt"
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
	loader := methodarea.NewClassLoader(cp)

	class := loader.LoadClass(c.classname)

	object := rt.NewObject(class)

	fmt.Println(object.Attributes)
}
