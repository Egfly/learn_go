package main

import (
	"fmt"
	"time"
)

func recursive(n int) {

	time.Sleep(time.Second)
	if n > 10 {
		return
	}
	fmt.Println(n)
	recursive(n + 1)
}

func main() {
	recursive(1)
}
