package main

import (
	"errors"
	"fmt"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test1() {
	err := initConfig()
	if err != nil {
		panic(err)
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	b := 0
	a := 100 / b
	fmt.Println(a)
}

func main() {
	var i int
	fmt.Println(i)

	// new()返回的是指针
	j := new(int)
	fmt.Println(j)
	*j = 100
	fmt.Println(*j)

	var a []int
	a = append(a, 10, 20, 30)
	fmt.Println(a)
	a = append(a, a...)
	fmt.Println(a)

	// defer
	test()

	// panic
	test1()
}
