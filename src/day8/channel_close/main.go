package main

import "fmt"

/**
channel 阻塞时，需要主动关闭channel，否则线程将会一直阻塞
close(channel)

如何检测channel是否关闭
variable， ok := <- channel
if ok == false {
	fmt.Println("channel is closed")
}
*/

func main() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	for {
		var b int
		b, ok := <-ch
		if ok == false {
			fmt.Println("channel is closed")
			break
		}
		fmt.Println(b)
	}
}
