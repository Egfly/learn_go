package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

/**
两次遍历hash表解法 O(n)
*/
func twoSum(nums []int, target int) []int {
	var result []int
	var m map[int]int
	if len(nums) < 2 {
		return result
	}
	m = make(map[int]int)
	for index, item := range nums {
		m[item] = index
	}

	for index, item := range nums {
		temp := target - item
		if v, ok := m[temp]; ok && v != index {
			result = append(result, index)
			result = append(result, v)
			break
		}
	}

	return result
}

/**
一次遍历hash表解法 O(n)
*/
func twoSumOneFor(nums []int, target int) []int {
	var result []int
	var m map[int]int
	if len(nums) < 2 {
		return result
	}
	m = make(map[int]int)
	for index, item := range nums {
		temp := target - item
		if v, ok := m[temp]; ok {
			result = append(result, v)
			result = append(result, index)
			break
		} else {
			m[item] = index
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSumOneFor(nums, target)
	fmt.Println(result)
}
