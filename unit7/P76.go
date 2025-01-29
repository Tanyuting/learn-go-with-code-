package main

import "fmt"

func main() {
	var arr [2][3]int16
	fmt.Println(arr)

	fmt.Printf("aar addres is : %p", &arr)
	fmt.Printf("aar [0] addres is:%p", &arr[0])
	fmt.Printf("aar[0] addres is: %p", &arr[0][0])

	fmt.Printf("aar [1] addres is:%p", &arr[1])
	fmt.Printf("aar[0][0] addres is: %p", &arr[1][0])

	arr[0][1] = 47
	arr[0][0] = 82
	arr[1][1] = 25
	fmt.Println(arr)

	var arr1 [2][3]int = [2][3]int{{1, 4, 7}, {2, 5, 8}}
	fmt.Println(arr1)
}
