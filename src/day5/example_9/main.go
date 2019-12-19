package main

import "fmt"

/**
实现String()
如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个
变量的String()进行输出。
*/
type car struct {
	weight int
	name   string
}

func (p car) run() {
	fmt.Println("running")
}

/**
字符化输出train
如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个
变量的String()进行输出。
*/
func (p train) String() string {
	str := fmt.Sprintf("train struct : [name:%s] [weight:%d]", p.c.name, p.c.weight)
	return str
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

	// 自定义string
	fmt.Printf("%s", b)
}
