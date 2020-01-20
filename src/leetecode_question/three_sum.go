package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i, v := range nums {
		temp := nums[i+1:]
		if i-1 >= 0 && nums[i] == nums[i-1] { //跳过重复值
			continue
		}
		l, r := 0, len(temp)-1
		for a := 0; a < len(temp); a++ {
			if l >= r {
				break
			}
			sum := v + temp[l] + temp[r]
			if sum > 0 {
				r--
			}
			if sum < 0 {

				l++
			}
			if sum == 0 {
				res = append(res, []int{v, temp[l], temp[r]})
				for l < r && temp[r] == temp[r-1] { //跳过重复值
					r--
				}
				for l < r && temp[l] == temp[l+1] { //跳过重复值
					l++
				}
				l++
				r--
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	var res [][]int
	var mid map[string]int
	var targetMap map[int]int
	mid = make(map[string]int)
	sort.Ints(nums)
	for i, v := range nums {
		if i >= 1 && v == nums[i-1] {
			continue
		}
		temp := nums[i+1:]
		targetMap = make(map[int]int)
		for _, item := range temp {
			if _, ok := targetMap[item]; ok {
				str := strconv.Itoa(v) + "," + strconv.Itoa(-v-item) + "," + strconv.Itoa(item)
				if _, ok := mid[str]; !ok {
					res = append(res, []int{v, -v - item, item})
					mid[str] = 1
				}
			} else {
				targetMap[-v-item] = 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}
	fmt.Println(threeSum2(nums))
	fmt.Println(threeSum(nums))

}
