package main

import (
	"fmt"
)

/**
channel写是阻塞的
如果定义一个容量为10的整形channel，在插入第11个数时，线程将会被阻塞，直到channel中有空闲位置时，从而将数据插入
*/
func calc(targetChan, sushuChan chan int, exitChan chan bool) {
	for v := range targetChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag == true {
			sushuChan <- v
		}
	}

	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	sushuChan := make(chan int, 1000)
	// 执行结束线程channel
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, sushuChan, exitChan)
	}

	go func() {
		// 等待所有计算的协程退出
		for i := 0; i < 8; i++ {
			<-exitChan
			fmt.Println("wait goroutine", i, "exited")
		}
		close(sushuChan)
	}()

	for v := range sushuChan {
		fmt.Println(v)
	}
}
