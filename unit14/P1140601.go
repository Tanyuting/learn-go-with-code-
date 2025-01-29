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
	fmt.Println("调用了Print()方法")
	fmt.Println("学生的名字：", s.Name)
}

func (s Student) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s *Student) Set(name string, age int) { // Use pointer receiver
	s.Name = name
	s.Age = age
}

func TestStudentStruct(a interface{}) {
	val := reflect.ValueOf(a)

	// Handle pointer case
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fmt.Println("结构体内容:", val)

	n1 := val.NumField()
	fmt.Println("字段数量:", n1)

	for i := 0; i < n1; i++ {
		fmt.Printf("第%d个字段的值：%v\n", i, val.Field(i))
	}

	fmt.Println()

	n2 := val.NumMethod()
	fmt.Println("方法数量:", n2)

	// Print available methods for debugging
	for i := 0; i < n2; i++ {
		fmt.Printf("方法 %d: %s\n", i, val.Type().Method(i).Name)
	}

	// Call a method safely if available
	if n2 > 0 {
		val.Method(0).Call(nil) // Call the first available method
	}
}

func main() {
	s := Student{Name: "张三", Age: 18}
	TestStudentStruct(s)
}
