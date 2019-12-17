package main

import "fmt"

func Adder() func(int) int {
	var x int
	return func(d int) int {
		fmt.Println(x)
		x += d
		return x
	}
}

func main() {
	f := Adder()
	f(1)
	f(100)
	f(1000)
	//fmt.Println(f(1))
	//fmt.Println(f(100))
	//fmt.Println(f(1000))
}
