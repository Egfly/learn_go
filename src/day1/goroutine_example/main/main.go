package main

import "fmt"
import "day1/goroutine_example/goroutine"

func main() {
	fmt.Println("goroutine start")
	pipe := make(chan int, 1)
	go goroutine.Add(300, 200, pipe)
	sum := <-pipe
	fmt.Println(sum)
}
