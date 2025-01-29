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

	num2 := 90 + reValue.Int()
	fmt.Println(num2)

	i2 := reValue.Interface()
	n := i2.(int)
	n2 := n + 12
	fmt.Println(n2)

}

func main() {

	var num int = 100
	testReflect(num)
}
