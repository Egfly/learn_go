package main

import "fmt"

/**
求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
*/
func sum(n int) int {
	s := 1
	sum := 0
	for i := 1; i <= n; i++ {
		s = s * i
		sum += s
	}

	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(sum(n))
}
