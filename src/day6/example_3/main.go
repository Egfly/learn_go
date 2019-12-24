package main

import "fmt"

/**
接口嵌套
*/
type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWrite interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("reading data....")
}

func (f *File) Write() {
	fmt.Println("writing data....")
}

func Test(rw ReadWrite) {
	rw.Read()
	rw.Write()
}

func main() {
	var f File
	Test(&f)
}
