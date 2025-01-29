package main

import "fmt"

func main() {
	//var aar1 = [3]int{3,6,9}
	//fmt.Println(aar1)
	//fmt.Printf("数组的类型：%T",aar1)

	var arr3 = [3]int{3, 6, 9}
	test1(&arr3)
	fmt.Println(arr3)

}

func test1(arr *[3]int) {

	(*arr)[0] = 7
}
