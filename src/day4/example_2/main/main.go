package main

import "fmt"

func test() {
	s1 := new([]int)
	//(*s1)[0] = 100 // panic: runtime error: index out of range 没有初始化
	*s1 = make([]int, 10)
	(*s1)[0] = 5
	fmt.Println(s1)
	s2 := make([]int, 10)
	s2[0] = 100
	fmt.Println(s2)
}

func main() {
	test()
}
