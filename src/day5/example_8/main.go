package main

import "fmt"

type car struct {
	weight int
	name   string
}

func (p car) run() {
	fmt.Println("running")
}

/**
bike 继承了car的属性 weight 和 name
*/
type bike struct {
	car
	lunzi int
}

type train struct {
	c car
}

func main() {
	// bike 继承了car的属性 weight 和 name 已经run方法
	var a bike
	a.name = "bike"
	a.weight = 100
	a.lunzi = 2
	fmt.Println(a)
	a.run()

	var b train
	b.c.weight = 200
	b.c.name = "train"
	b.c.run()
}
