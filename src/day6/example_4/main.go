package main

import "fmt"

func Test(a interface{}) {
	b, ok := a.(int)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	b += 3
	fmt.Println(b)
}

func just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params type is bool, value is %v \n", index, v)
		case int, int32, int64:
			fmt.Printf("%d params type is int, value is %v \n", index, v)
		case float32, float64:
			fmt.Printf("%d params type is float, value is %v \n", index, v)
		case string:
			fmt.Printf("%d params type is string, value is %v \n", index, v)
		}
	}
}

/**
类型断言
*/
func main() {
	var a interface{}
	var b int
	a = b
	c := a.(int)
	fmt.Printf("%T %T %T\n", a, b, c)

	str := "This is a string"
	Test(str)

	just(a, str)
}
