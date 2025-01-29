package main

import "fmt"

type AInterface interface {
	a()
}

type BInterface interface {
	b()
}

type Stu struct {
}

func (s Stu) a() {
	fmt.Println("aaaa")

}

func (s Stu) b() {

	fmt.Println("bbb")

}

func main() {

	var s Stu
	var a AInterface = s
	var b BInterface = s
	a.a()
	b.b()
}
