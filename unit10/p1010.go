package main

import "fmt"

type Student struct {
	Name string
}

//方法

func (s Student) test01() {

	fmt.Println(s.Name)
}

//函数

func method01(s Student) {
	fmt.Println(s.Name)

}

func main() {
	var s Student = Student{"lili"}
	method01(s)

	s.test01()

}
