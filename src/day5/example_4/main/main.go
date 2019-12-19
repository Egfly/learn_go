package main

import (
	"encoding/json"
	"fmt"
)

/**
struct tag
*/

/**
若不使用tag, struct内的元素名必须首字母大写才能对外可见
才能被转成json
*/
type Student1 struct {
	Name  string
	Age   int
	Score float32
}

/**
使用tag
*/
type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
}

func main() {
	//var stu Student1 = Student1{
	//	Name:  "stu01",
	//	Age:   18,
	//	Score: 80,
	//}
	//data, err := json.Marshal(stu)
	//if err != nil {
	//	fmt.Println("json encode stu failed, err:", err)
	//}
	//fmt.Println(string(data))

	var stu Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 80,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
	}
	fmt.Println(string(data))
}
