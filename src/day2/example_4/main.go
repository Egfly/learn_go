package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	second := time.Now().Unix()
	result := second % Female
	fmt.Println("当前秒数：", second, "除以Female的结果：", result)
	if result == 0 {
		fmt.Println(Female)
	} else {
		fmt.Println(Man)
	}
}
