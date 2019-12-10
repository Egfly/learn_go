package main

import "fmt"

/**
斐波那契数列解法
*/

// 递归解法
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 动态规划解法
func dynamicFib(n int, memory [10]int) int {
	if n <= 1 {
		return n
	} else if memory[n] == 0 {
		memory[n] = fib(n-1) + fib(n-2)
	}
	return memory[n]
}

func main() {
	nums := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("递归解法....")
	for _, item := range nums {
		fmt.Println(fib(item))
	}
	fmt.Println("动态规划解法....")
	for _, item := range nums {
		var memory [10]int
		fmt.Println(dynamicFib(item, memory))
	}
}
