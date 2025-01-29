package main

import "fmt"

func cal(num1 int, num2 int) int {
	var sum int = 0
	sum += num1
	sum += num2
	return sum

}

func main() {
	sum := cal(10, 20)
	fmt.Println(sum)

	var num3 int = 30
	var num4 int = 50
	sum1 := cal(num3, num4)
	fmt.Println(sum1)

}
