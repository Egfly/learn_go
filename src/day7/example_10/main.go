package main

import (
	"fmt"
	"os"
)

/**
自定义错误
*/
type PathError struct {
	Op   string
	Path string
	err  string
}

func (p *PathError) Error() string {
	return p.Op + "" + p.Path + ":" + p.err
}

func OpenFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &PathError{
			Op:   "read",
			Path: fileName,
			err:  err.Error(),
		}
	}
	defer file.Close()

	return nil
}

func main() {
	err := OpenFile("C:/ADSADASD")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error", v)
	default:
		fmt.Println(1111)
	}
}
