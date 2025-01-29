package main

import "fmt"

func getSum() func(int) int {

	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(0))
	fmt.Println(f(3))

	fmt.Println("----------------")
	fmt.Println(getSum01(0, 1))
	fmt.Println(getSum01(1, 2))

}

func getSum01(sum int, num int) int {
	//var sum int = 0
	sum = sum + num
	return sum
}

//不使用闭包的时候，我想保留的值，不可以反复用
//闭包场景： 可以保留一次数值后反复使用
