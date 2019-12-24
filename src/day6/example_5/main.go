package main

import "fmt"

/**
实现单链表
*/

type Node struct {
	data interface{}
	next *Node
}
type Link struct {
	head *Node
	tail *Node
}

/**
头部插入
*/
func (p *Link) InsertHead(data interface{}) {
	node := &Node{
		data: data,
		next: nil,
	}
	if p.head == nil {
		p.head = node
		p.tail = node
	} else {
		node.next = p.head
		p.head = node
	}
}

/**
尾部插入
*/
func (p *Link) InsertTail(data interface{}) {
	node := &Node{
		data: data,
		next: nil,
	}
	if p.tail == nil {
		p.head = node
		p.tail = node
	} else {
		p.tail.next = node
		p.tail = node
	}
}

func (p *Link) Trans() {
	head := p.head
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func main() {
	var intLink Link
	for i := 0; i < 10; i++ {
		intLink.InsertTail(i)
	}

	intLink.Trans()
}
