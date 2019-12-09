package main

import "fmt"
import "day1/package_example/calc"

func main() {
	a, b := 700, 300
	fmt.Println(calc.Add(a, b))
	fmt.Println(calc.Sub(a, b))
}
