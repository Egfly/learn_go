package main

import "fmt"

/**

 */

func main() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	//ch <- 100 // panic: send on closed channel
	for v := range ch {
		fmt.Println(v)
	}

}
