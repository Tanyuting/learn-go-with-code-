package main

import "fmt"

func main() {
	var age int = 18
	fmt.Println("age is the memory address :", &age)

	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr is the :", *ptr)
}
