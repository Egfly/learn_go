package main

import (
	"time"
)

func main() {
	//pipe = make(chan int, 1)
	//println("out start")
	//go add(300, 200)
	//println("out end")
	//// 协程中存放数据到channel 主线程接收channel数据时
	//// 如果channel里没有数据，主线程会阻塞，等待协程执行结束，然后继续往下执行
	//sum := <-pipe
	//println(sum)
	//time.Sleep(3 * time.Second)
	//println(sum)
	//time.Sleep(5 * time.Second)
	//println(sum)
	for i := 0 ; i < 100; i++ {
		go testRoute(i)
	}
	//time.Sleep(2)
	//fmt.Println("start")
	//testPip()
	//fmt.Println("end")
	time.Sleep(10 * time.Second)
}
