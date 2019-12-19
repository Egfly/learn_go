package main

import "fmt"

/**
任意类型通过type设置别名后，是一个新的类型。与原类型不同
*/
type Integer int //alias 设置int的别名 Integer

func main() {
	var a Integer = 100
	var b int = 10
	a = b // 不同类型无法赋值
	a = Integer(b)
	fmt.Println(a, b)
}
