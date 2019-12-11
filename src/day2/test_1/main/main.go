package main

import "fmt"

/**
判断 101-200 之间有多少个素数，并输出所有素数
*/
func main() {
	sum := 0
	for i := 101; i < 201; i++ {
		isSu := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isSu = false
				break
			}
		}
		if isSu == true {
			fmt.Println(i)
			sum++
		}
	}

	fmt.Println("总数：", sum)
}
