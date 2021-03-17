package main

import "./classpath"
import "./methodarea"

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

	loader.LoadClass(c.classname)
}
