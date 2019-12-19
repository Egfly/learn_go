package main

import "fmt"

/**
冒泡排序
*/
func mpSort(a []int) {
	length := len(a)
	if length == 0 {
		return
	}
	index := 0
	for index < length {
		for i := 0; i < length; i++ {
			if (length-i-2 >= 0) && a[length-1-i] < a[length-2-i] {
				a[length-i-1], a[length-i-2] = a[length-i-2], a[length-i-1]
			}
		}
		index++
	}

}

func main() {
	a := []int{5, 7, 1, 4, 9, 10, 8}
	fmt.Println(a)
	mpSort(a)
	fmt.Println(a)
}
