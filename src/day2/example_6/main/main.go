package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成10个随机整数
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	// 生成10个小于100的随机整数
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	// 生成10个随机浮点数
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}
