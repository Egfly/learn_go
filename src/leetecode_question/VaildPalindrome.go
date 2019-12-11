package main

import (
	"fmt"
	"strings"
)

/**
leetcode 125、
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/

/**
判断是否为回文字符串
双指针解法：O(n)
*/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 处理字符串中的空格
	s = strings.Replace(s, " ", "", -1)
	// 字符串处理成小写
	s = strings.ToLower(s)
	end, start, result := len(s)-1, 0, true
	//验证是否为回文字符串
	for start < end {
		if !((s[start] >= 'a' && s[start] <= 'z') || (s[start] >= '0' && s[start] <= '9')) {
			start++
			continue
		}
		if !((s[end] >= 'a' && s[end] <= 'z') || (s[end] >= '0' && s[end] <= '9')) {
			end--
			continue
		}
		if s[start] != s[end] {
			result = false
			break
		}
		start++
		end--
	}

	return result
}

func main() {
	s := "A man, a plan, a canal: Panama"
	result := isPalindrome(s)
	fmt.Println(result)
}
