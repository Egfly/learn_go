package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	fmt.Println(right)
	max := 0
	for {
		if left >= right {
			break
		}
		min := height[left]
		if height[left] > height[right] {
			min = height[right]
		}
		temp := (right - left + 1) * min
		if max < temp {
			max = temp
		}
		fmt.Println("left= ", left, " right= ", right)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return max
}
func main() {
	data := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(data))
}
