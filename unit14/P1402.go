package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {

	reType := reflect.TypeOf(i)
	fmt.Println("reType:", reType)

	reValue := reflect.ValueOf(i)
	fmt.Println("reValue:", reValue)
	fmt.Printf("reType 的类型：%T \n", reType)

	i2 := reValue.Interface()
	n, flag := i2.(Student)
	if flag == true {

		fmt.Printf("student name:%v, age： %v", n.Name, n.Age)
	}

}

type Student struct {
	Name string
	Age  int
}

func main() {

	stu := Student{
		Name: "tantan",
		Age:  18,
	}
	testReflect(stu)
}
