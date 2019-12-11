package main

import "fmt"

/**
一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
编程找出1000以内的所有完数。
*/
func main() {
	for i := 0; i < 1000; i++ {
		sum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += i / j
			}
		}
		if sum == i {
			fmt.Println(i)
		}
	}
}
