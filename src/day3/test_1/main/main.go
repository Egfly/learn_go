package main

import "fmt"

/**
打印九九乘法表
*/
func table() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d   ", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	table()
}
