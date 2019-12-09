package main

func testPip() {
	pip := make(chan int, 3)
	pip <- 1
	pip <- 2
	pip <- 3

	//var a int
	//a =<- pip
	//sum = <- pip
	//pip <- 4

	//println(a)
	//println("sum=", sum)
	println(len(pip))
}
