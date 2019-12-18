package main

import (
	"fmt"
	"math/rand"
)

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

//遍历链表
func foreach(p *student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func tailInsert(p *student) {
	// 循环插入 生成10个节点的链表
	for i := 0; i < 10; i++ {
		var students = student{
			name:  fmt.Sprintf("student%d", i+3),
			age:   i + 3,
			score: rand.Float32() * 100,
		}
		p.next = &students
		p = &students
	}
}

func main() {
	head := student{
		name:  "head",
		age:   1,
		score: 1,
		next:  nil,
	}
	// 尾部插入
	student1 := student{
		name:  "student1",
		age:   2,
		score: 2,
	}
	head.next = &student1
	// 尾部插入
	student2 := student{
		name:  "student2",
		age:   3,
		score: 3,
	}
	student1.next = &student2
	foreach(&head)
	// 尾部插入10个元素
	tailInsert(&student2)
	foreach(&head)
}
