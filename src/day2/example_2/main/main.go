package main

import (
	sum "day2/example_2/add"
	"fmt"
)

/**
test2. 包别名的应用，开发一个程序，使用包别名来访问包中的函数？
*/
func main() {
	a, b := 3, 4
	fmt.Printf("%d+%d=%d", a, b, sum.Add(a, b))
}
