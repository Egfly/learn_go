package main

import (
	"fmt"
	"os"
)

/**
命令行参数读取：os.Args
*/
func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))
	for index, v := range os.Args {
		fmt.Printf("key:%d  value:%s\n", index, v)
	}
}
