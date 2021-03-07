package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildCardEntry(path string) Entry {
	dirPath := path[:len(path)-1]

	var compositeEntry CompositeEntry

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != dirPath {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path,
			".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path,
			".ZIP") {
			entry := newZipEntry(path)
			compositeEntry = append(compositeEntry, entry)
		}

		return nil
	})

	return compositeEntry
}
