package main

import "fmt"

type Student struct {
	Name string
}

func method01(s Student) {
	fmt.Println(s.Name)
}

func method02(s *Student) {

	fmt.Println((*s).Name)

}

func main() {
	var s Student = Student{"lili"}

	method01(s)

	method02(&s)
}
