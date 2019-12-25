package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("E:/go_projects/learn_go/src/day7/example_4/test.txt")
	if err != nil {
		fmt.Println("Open file failed, err:", err)
		return
	}
	defer file.Close()

}
