package main

import "fmt"

func main() {
	var intarr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var slice []int = intarr[1:4]
	fmt.Println(len(slice))

	slice2 := append(slice, 88, 50)
	fmt.Println(slice2)
	fmt.Println(slice)

	slice3 := []int{99, 100}
	slice = append(slice, slice3...)
	fmt.Println(slice)

}
