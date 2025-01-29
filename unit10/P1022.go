package main

import "fmt"

type Sayhello interface {
	Sayhello()
}

type Chinese struct {
}

func (person Chinese) Sayhello() {
	fmt.Println("你好")
}

type American struct {
}

func (Person American) Sayhello() {
	fmt.Println("hi")
}

func greet(s Sayhello) {
	s.Sayhello()
}

type integer int

func (i integer) Sayhello() {
	fmt.Println("say hi +", i)
}

func main() {

	var i integer = 20
	var s Sayhello = i
	s.Sayhello()

	c := Chinese{}
	a := American{}

	greet(a)
	greet(c)

	//var s Sayhello = c
	//s.Sayhello()

}
