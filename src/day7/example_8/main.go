package main

import (
	"encoding/json"
	"fmt"
)

/**
Json数据协议

*/

type User struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func testStruct() {
	user1 := User{
		UserName: "user1",
		NickName: "用户1",
		Age:      27,
		Sex:      "男",
		Email:    "1147732@qq.com",
		Phone:    "110",
	}

	// 返回的是[]byte
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testInt() {
	var age int = 100

	// 返回的是[]byte
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user_map"
	m["age"] = 10
	m["sex"] = "女"

	// 返回的是[]byte
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testSlice() {
	var s []map[string]interface{}
	m := make(map[string]interface{})
	m["username"] = "user_map"
	m["age"] = 10
	m["sex"] = "女"
	s = append(s, m)

	u1 := make(map[string]interface{})
	u1["username"] = "user_slice"
	u1["age"] = 20
	u1["sex"] = "男"
	s = append(s, u1)

	// 返回的是[]byte
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func main() {
	//testStruct()
	//testInt()
	//testMap()
	testSlice()
}
