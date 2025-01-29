package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	var s1 Student = Student{"xiao li ", 19}

	fmt.Println(s1)

	var s2 Student = Student{
		Name: "lili",
		Age:  19,
	}
	fmt.Println(s2)

	var s3 *Student = &Student{"mingming", 19}
	fmt.Println(*s3)
}
