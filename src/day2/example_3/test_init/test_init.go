package test_init

import (
	// 如果只想导入包，但是不使用，可以使用 _ 做为别名
	_ "day2/example_3/test_init_2"
	"fmt"
)

var Name string = "xxxxxxx"
var Age int = 0
var num1 int = 1
var num2 int = 2

func init() {
	fmt.Printf("init function runs after global variable assignment")
	fmt.Println()
	fmt.Printf("Name= %s, Age= %d", Name, Age)
	fmt.Println()
	fmt.Println("num1=", num1, "num2=", num2)
	Name = "Eric"
	Age = 26
	num1 = 20
	num2 = 30
}

func Test() {
	fmt.Println("num1=", num1, "num2=", num2)
}
