package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Id   string
	Age  int
}

type Book struct {
	Name   string
	Author string
	Id     int
}
type StudentArr []Student

func (p StudentArr) Len() int {
	return len(p)
}

func (p StudentArr) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p StudentArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var stus StudentArr
	for i := 0; i < 10; i++ {
		item := Student{
			Name: fmt.Sprintf("student%d", rand.Intn(100)),
			Id:   fmt.Sprintf("110%d", rand.Int()),
			Age:  rand.Intn(100),
		}
		stus = append(stus, item)
	}

	for _, v := range stus {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	sort.Sort(stus)

	for _, v := range stus {
		fmt.Println(v)
	}
}
