package main

import "fmt"

func main() {

	var intarr [6]int = [6]int{1, 3, 4, 5, 6, 7}
	var slice []int = intarr[1:4]
	//fmt.Println(slice[0])
	//fmt.Println(slice[1])
	//fmt.Println(slice[3])
	slice2 := slice[1:2]
	fmt.Println(slice2)
	slice2[0] = 66
	fmt.Println(intarr)
	fmt.Println(slice)

}
