package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	str := fmt.Sprintf("name = %v, age = %v", s.Name, s.Age)
	return str
}

func main() {
	stu := Student{
		Name: "lili",
		Age:  20,
	}
	fmt.Println(&stu)
}
