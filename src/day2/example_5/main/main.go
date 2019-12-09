package main

import "fmt"

func modify(a int) {
	a = 10
}

func modify1(a *int) {
	*a = 20
}
func main() {
	a := 5
	b := make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	// 值类型传递时是复制一份数据，传递出去。修改传递出去的值，不影响本身
	modify(a)
	fmt.Println("a=", a)

	// 指针传递：指针指向存放数据的内存空间地址，修改内存空间中的数据后，其他指向这个内存空间的变量的值一起改变
	modify1(&a)
	fmt.Println("a=", a)
}
