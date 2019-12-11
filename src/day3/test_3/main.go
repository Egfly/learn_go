package main

import "fmt"

/**
输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
左读完全相同的字符串。
*/
func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	result := true
	for {
		if start >= end {
			break
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
	var s string
	fmt.Println("请输入字符串")
	fmt.Scanf("%s/n", &s)
	result := isPalindrome(s)
	fmt.Println(result)
}
