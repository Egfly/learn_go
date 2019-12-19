package main

import (
	"fmt"
	"time"
)

type Cart struct {
	name string
	age  int
}

/**
匿名字段
*/
type Train struct {
	Cart
	int
	start time.Time
}

func main() {
	var t Train
	t.name = "train"
	t.age = 100
	t.int = 200
	fmt.Println(t)

	t.Cart.name = "001"
	t.Cart.age = 300
	fmt.Println(t)
}
