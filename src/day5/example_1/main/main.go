package main

import "fmt"

/**
1. 用来自定义复杂数据结构
2. struct里面可以包含多个字段（属性）
3. struct类型可以定义方法，注意和函数的区分
4. struct类型是值类型
5. struct类型可以嵌套
6. Go语言没有class类型，只有struct类型
*/

type student struct {
	Name  string
	Age   int32
	score float32
}

func testStruct() {
	var stu student
	stu.Age = 18
	stu.Name = "egfly"
	stu.score = 100
	fmt.Println(stu)
	fmt.Printf("Name's address:%p\n", &stu.Name)
	fmt.Printf("Age's address:%p\n", &stu.Age)
	fmt.Printf("score's address:%p\n", &stu.score)
	fmt.Println()
	var stu1 *student = &student{
		Name: "test1",
		Age:  20,
	}
	stu1.score = 150
	fmt.Printf("Name's address:%p\n", &stu1.Name)
	fmt.Printf("Age's address:%p\n", &stu1.Age)
	fmt.Printf("score's address:%p\n", &stu1.score)
	fmt.Println()
	var stu2 = student{
		Name:  "test1",
		Age:   20,
		score: 110,
	}
	fmt.Printf("Name's address:%p\n", &stu2.Name)
	fmt.Printf("Age's address:%p\n", &stu2.Age)
	fmt.Printf("score's address:%p\n", &stu2.score)

}

func main() {
	testStruct()
}
