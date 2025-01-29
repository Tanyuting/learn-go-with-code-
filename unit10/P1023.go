package main

import "fmt"

type CInterface interface {
	c()
}

type BInterface interface {
	b()
}

type AInterface interface {
	BInterface
	CInterface
	a()
}

type Stu struct {
}

func (s Stu) a() {
	fmt.Println("a")
}

func (s Stu) b() {
	fmt.Println("b")
}

func (s Stu) c() {
	fmt.Println("c")
}

func main() {
	var s Stu
	var a AInterface = s
	a.a()
	a.b()
	a.c()
}
