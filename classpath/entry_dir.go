package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	path string
}

func newDirEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &DirEntry{
		path: absPath,
	}
}

func (entry *DirEntry) ReadClass(classFile string) ([]byte, Entry, error) {
	// 从当前目录下读取指定的class文件

	filePath := filepath.Join(entry.path, classFile)

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, nil, err

	}

	return content, entry, nil
}

func (entry DirEntry) string() string {
	return "DirEntry path = " + entry.path
}
