package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry {
	paths := strings.Split(path, FileSeparator)

	entrys := make([]Entry, 0, len(paths))

	for _, path := range paths {
		entrys = append(entrys, NewEntry(path))
	}

	return entrys
}

func (entry CompositeEntry) string() string {
	var strList []string

	for _, e := range entry {
		strList = append(strList, e.string())
	}

	return strings.Join(strList, ";")
}

func (entry CompositeEntry) ReadClass(classFile string) ([]byte, Entry, error) {
	for _, item := range entry {
		content, entry, err := item.ReadClass(classFile)
		if err == nil {
			return content, entry, nil
		}
	}

	return nil, nil, errors.New("cant find this " + classFile)
}
