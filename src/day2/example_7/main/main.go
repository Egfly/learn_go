package main

import (
	re "day2/example_7/reverse"
	"fmt"
)

func main() {
	// 字符串拼接 +
	//str1, str2 := "中华人民", "共和国"
	str1, str2 := "hello", "world"
	//str3 := str1 + " " + str2
	// 字符串拼接， 格式化写入
	str3 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str3)
	// 获取字符串长度
	strLen := len(str3)
	fmt.Printf("len(str3)=%d\n", strLen)

	//截取字符串
	substr := str3[0:5]
	substr = str3[5:]
	fmt.Printf("str3[0:5]=%s\n", substr)

	// 反转字符串
	fmt.Printf("reverse string：%s \n", re.Reverse(str3))
	fmt.Printf("reverse string：%s \n", re.Reverse1(str3))
}
