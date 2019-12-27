package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int64]int64)
	lock sync.Mutex
)

type task struct {
	n int64
}

func calc(t *task) {
	var sum int64
	sum = 1
	for i := int64(1); i < t.n; i++ {
		sum *= i
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i := int64(0); i < 16; i++ {
		t := &task{n: i}
		go calc(t)
	}

	time.Sleep(10 * time.Second)
	lock.Lock()
	for k, v := range m {
		fmt.Printf("%d = %d \n", k, v)
	}
	lock.Unlock()
}
