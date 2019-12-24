package main

import (
	"fmt"
	"reflect"
)

/**
反射：可以在运行时动态获取变量的相关信息
*/

type Student struct {
	Name string
	Age  int
}

func test(b interface{}) {
	fmt.Println(reflect.TypeOf(b))
	v := reflect.ValueOf(b)
	fmt.Println(v.Kind())
	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	c := val.Int()
	fmt.Println(c)
}
func main() {
	//a := 20
	//test(a)

	b := Student{
		Name: "student",
		Age:  27,
	}
	test(b)

	c := 30
	testInt(c)
}
