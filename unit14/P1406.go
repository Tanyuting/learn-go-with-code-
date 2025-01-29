package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) Print() {
	fmt.Println("调用了Pring()方法")
	fmt.Println("学生的名字：", s.Name)
}

func (s Student) GetSum(n1, n2 int) int {

	return n1 + n2
}

func (s *Student) Set(name string, age int) {
	s.Name = name
	s.Age = age

}

func TestStudentStruct(a interface{}) {
	val := reflect.ValueOf(a)
	fmt.Println(val)

	n1 := val.NumField()
	fmt.Println(n1)

	for i := 0; i < n1; i++ {
		fmt.Printf("第%d个字段的值：%v", i, val.Field(i))
	}
	fmt.Println()

	n2 := val.NumMethod()
	fmt.Println(n2)

	val.Method(2).Call(nil)

}

func main() {

	s := Student{
		Name: "taannn",
		Age:  18,
	}
	TestStudentStruct(s)
}
