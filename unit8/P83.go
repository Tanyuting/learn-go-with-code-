package main

import "fmt"

func main() {

	slice := make([]int, 4, 20)
	slice[0] = 66
	slice[1] = 77
	slice[2] = 88
	slice[3] = 99
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v \t", i, slice[i])
	}
	fmt.Println("-------------\n")
	for i, v := range slice {
		fmt.Printf("下标i:%v,元素 : %v", i, slice[i])
	}
}
