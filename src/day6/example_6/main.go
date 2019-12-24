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

func testStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect a struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Field(i))
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod)

	var param []reflect.Value
	val.Method(0).Call(param)
}

func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func (s Student) Print() {
	fmt.Println(s)
}

/**
当传入的参数是指针时，需要使用Elem()方法才能调用struct的type和valueOf
*/
func testReflectPtr(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Elem().Kind()
	if kd != reflect.Struct {
		fmt.Println("expect a struct")
		return
	}

	num := val.Elem().NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i))
	}

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod)

	var param []reflect.Value
	val.Elem().Method(0).Call(param)
}
func main() {
	//a := 20
	//test(a)

	//b := Student{
	//	Name: "student",
	//	Age:  27,
	//}
	//test(b)
	//
	//c := 30
	//testInt(c)

	//testStruct(123)

	d := Student{
		Name: "stu01",
		Age:  20,
	}
	//testStruct(d)

	testReflectPtr(&d)
}
