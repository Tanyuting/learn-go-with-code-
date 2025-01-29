package main

import "fmt"

var Func01 = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	//定义匿名函数
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println(result)

	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}
	result01 := sub(30, 70)
	fmt.Println(result01)

	reuslt03 := Func01(3, 4)
	fmt.Println(reuslt03)

}
