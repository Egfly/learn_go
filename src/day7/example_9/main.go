package main

import (
	"encoding/json"
	"fmt"
)

/**
反序列化json字符串
*/
type User struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func testStruct() (result string, err error) {
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
		err = fmt.Errorf("json.marshal failed, err:", err)
		return
	}
	result = string(data)
	return
}

/**
反序列化json
*/
func test() {
	jsonStr, err := testStruct()
	if err != nil {
		fmt.Println("test struct failed", err)
		return
	}
	var user1 User
	err1 := json.Unmarshal([]byte(jsonStr), &user1)
	if err1 != nil {
		fmt.Println("unmarshal json string failed! err:", err1)
		return
	}
	fmt.Println(user1)
}

func testMap() (result string, err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user_map"
	m["age"] = 10
	m["sex"] = "女"

	// 返回的是[]byte
	data, err := json.Marshal(m)
	if err != nil {
		err = fmt.Errorf("json.marshal failed, err:", err)
		return
	}
	result = string(data)
	return
}

func unMarshalMap() {
	str, err := testMap()
	if err != nil {
		fmt.Println("test map failed! err:", err)
		return
	}

	var m map[string]interface{}
	m = make(map[string]interface{})
	err = json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("unmarshal map failed! err:", err)
		return
	}

	fmt.Println(m)
}
func main() {
	//test()
	unMarshalMap()
}
