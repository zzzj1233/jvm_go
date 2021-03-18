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

func TestRt(t *testing.T) {
	path := "D:\\java\\jdk\\jre\\lib\\rt.jar"

	p := "java/lang/Object.class"

	reader, _ := zip.OpenReader(path)

	for _, f := range reader.Reader.File {
		if f.Name == p {
			fmt.Println("find ~")
		}
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
