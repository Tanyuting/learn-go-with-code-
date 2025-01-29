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

func main() {

	c := Chinese{}
	a := American{}

	greet(a)
	greet(c)

}
