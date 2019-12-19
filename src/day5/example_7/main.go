package main

import "fmt"

/**
Golang中的方法是作用在特定类型的变量上，因此自定义类型，都可以
有方法，而不仅仅是struct

定义：func (recevier type) methodName(参数列表)(返回值列表){}
*/

type integer int

func (p integer) print() {
	fmt.Println(p)
}

/**
struct是值类型，在传递时创建副本使用，要修改则要使用指针
*/
func (p *integer) set(value integer) {
	*p = value
}

type student struct {
	Name string
	Age  int
}

/**
struct是值类型，在传递时创建副本使用，要修改则要使用指针
*/
func (p *student) init(name string, age int) {
	p.Age = age
	p.Name = name
	fmt.Println(p)
}

func (p student) get() student {
	return p
}

type sl map[int]string

func (p *sl) set(v sl) {
	*p = v
}

func (p sl) print() {
	fmt.Println(p)
}
func main() {
	var stu student
	stu.init("test_method", 1)
	fmt.Println(stu)
	stu1 := stu.get()
	fmt.Println(stu1)

	var a integer = 100
	a.print()
	a.set(1000)
	a.print()

	var m sl = make(sl)
	m[0] = "a"
	m[1] = "b"
	m[2] = "c"
	m[3] = "d"
	var s sl = make(sl)
	fmt.Println(s)
	s.set(m)
	s.print()
}
