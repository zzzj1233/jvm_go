package classpath

import (
	"archive/zip"
	"bufio"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	path string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &ZipEntry{
		path: absPath,
	}
}

func (entry *ZipEntry) string() string {
	return "ZipEntry : absPath = " + entry.path
}

func (entry *ZipEntry) ReadClass(classFile string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(entry.path)

	if err != nil {
		return nil, nil, err
	}

	defer reader.Close()

	for _, file := range reader.File {
		if file.Name == classFile {
			f, err := file.Open()

			if err != nil {
				return nil, nil, err
			}

			defer f.Close()

			reader := bufio.NewReader(f)

			content, err := ioutil.ReadAll(reader)

			if err != nil {
				return nil, nil, err
			}

			return content, entry, nil
		}
	}

	return nil, nil, errors.New("cant find this " + classFile + " in " + entry.path)
}
