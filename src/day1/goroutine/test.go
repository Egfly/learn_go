package main

import (
	"fmt"
	"time"
)

var pipe chan int
func testRoute(i int) {
	fmt.Println(i)
}

func add(a int, b int) {
	fmt.Println("tread start")
	var sum int
	sum = a + b
	time.Sleep(5 * time.Second)
	pipe <- sum
	fmt.Println("tread end")
}
