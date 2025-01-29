package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {

	reType := reflect.TypeOf(i)

	reValue := reflect.ValueOf(i)

	k1 := reType.Kind()
	fmt.Println(k1)

	k2 := reValue.Kind()
	fmt.Println(k2)

	i2 := reValue.Interface()
	n, flag := i2.(Student)
	if flag == true {

		fmt.Printf("Struct's type: %T", n)
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
