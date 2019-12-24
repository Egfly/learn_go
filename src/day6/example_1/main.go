package main

import "fmt"

type Test interface {
}

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running..\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s di di di..\n", p.Name)
}

func main() {
	var a interface{}
	var b int
	a = b
	fmt.Printf("type of a %T\n", a)
	fmt.Printf("type of a %T\n", b)

	var car1 Car
	var bmw BMW
	bmw.Name = "宝马"
	car1 = &bmw
	fmt.Printf("car name is %s\n", car1.GetName())
	car1.DiDi()
	car1.Run()
	fmt.Printf("type of a %T\n", car1)

	var car2 Car
	car2 = &BMW{
		Name: "宝马mini",
	}
	fmt.Printf("car name is %s\n", car2.GetName())
	car2.DiDi()
	car2.Run()
}
