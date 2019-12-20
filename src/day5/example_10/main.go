package main

import "fmt"

/**
接口
接口定义使用 interface关键字
接口内只能包含方法
若要实现一个接口，必须实现接口内所有函数
*/
type Test interface {
	Print()
	Sleep()
}

type student struct {
	name  string
	age   int
	score int
}

func (p *student) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("score:", p.score)
}

func (p student) Sleep() {
	fmt.Println("student sleeping......")
}

type People struct {
	name string
	age  int
}

func (p People) Print() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
}

func (p People) Sleep() {
	fmt.Println("people sleeping......")
}

func main() {
	var t Test
	var stu student = student{
		name:  "stu1",
		age:   18,
		score: 100,
	}
	t = &stu
	fmt.Println(t)
	t.Print()
	t.Sleep()

	var pe People = People{
		name: "people",
		age:  20,
	}
	var t1 Test = pe
	t1.Sleep()
	t1.Print()
	fmt.Println(t1)
}
