package classpath

import (
	"os"
	"strings"
)

const FileSeparator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(classFile string) ([]byte, Entry, error)

	string() string
}

func NewEntry(path string) Entry {
	// 1. jar
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") ||
		strings.HasSuffix(path, "zip") || strings.HasSuffix(path, "ZIP") {
		return newZipEntry("")
		// 2. xx;xx
	} else if strings.Contains(path, FileSeparator) {
		return newZipEntry(path)

		// 2. xx;xx
	} else if strings.Contains(path, FileSeparator) {
		return newCompositeEntry(path)
		// 3. *
	} else if strings.HasSuffix(path, "*") {
		return newWildCardEntry(path)
		// 4. dir
	} else {
		return newDirEntry(path)
	}
}
