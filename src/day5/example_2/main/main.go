package main

import "fmt"

/**
结构体实现单链表
单链表：｛
	content //本身的内容
	next_ptr //下个元素的指针，不存在时是nil
｝
*/
type student struct {
	name  string
	age   int
	score float32
	next  *student
}

func main() {
	head := student{
		name:  "head",
		age:   1,
		score: 1,
		next:  nil,
	}

	student1 := student{
		name:  "student1",
		age:   2,
		score: 2,
	}
	head.next = &student1
	//遍历链表
	var p *student = &head
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}
