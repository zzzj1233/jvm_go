package test

import (
	"archive/zip"
	"fmt"
	"path/filepath"
	"testing"
)

func TestZip(t *testing.T) {

	const path = "C:\\Users\\zzzj\\Desktop\\img\\img.zip"

	reader, _ := zip.OpenReader(path)

	defer reader.Close()

	for _, i := range reader.File {
		fmt.Println(i.Name)
	}

}

func TestPath(t *testing.T) {

	const path = "../img*"

	fmt.Println(filepath.Abs(path))
}

type TestInter interface {
}

func TestNil(t *testing.T) {
	var a TestInter
	var b []string

	println(a == nil)
	println(b == nil)
}
