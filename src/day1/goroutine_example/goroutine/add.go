package goroutine

func Add(a int, b int, pipe chan int) {
	pipe <- a + b
}
