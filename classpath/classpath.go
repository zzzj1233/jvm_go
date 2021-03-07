package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	BootClassPath Entry
	ExtClassPath  Entry
	UsrClassPath  Entry
}

func Parse(jreOption string, cpOption string) *ClassPath {
	cp := &ClassPath{}

	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUsrClassPath(cpOption)

	return cp
}

func (this *ClassPath) parseBootAndExtClassPath(jreOption string) {
	jrePath := parseJrePath(jreOption)

	// boot 位于 lib
	this.BootClassPath = NewEntry(filepath.Join(jrePath, "lib", "*"))
	// ext  位于 lib/ext
	this.ExtClassPath = NewEntry(filepath.Join(jrePath, "lib", "ext", "*"))
}

func (this *ClassPath) parseUsrClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	this.UsrClassPath = NewEntry(cpOption)
}

func parseJrePath(jreOption string) string {
	// 1. 用户指定目录
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 2. 当前目录/jre

	if exists("./jre") {
		return "./jre"
	}
	// 3. JAVA_HOME路径/jre

	if jh := os.Getenv("JAVA_HOME"); jh != "" && exists(filepath.Join(jh, "jre")) {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
