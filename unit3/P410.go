package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Println(sum)
		if sum >= 300 {
			break
		}

	}
	fmt.Println("dddddd")
}
