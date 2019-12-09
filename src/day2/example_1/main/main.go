package main

import (
	"day2/example_1/add"
	"fmt"
)

/**
test1:一个程序包含两个包add和main，其中add包中有两个变量：Name和age。请问main
包中如何访问Name和age？
*/
func main() {
	fmt.Println(add.Name, "的年龄是", add.Age)
}
