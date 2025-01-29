package main

import (
	"fmt"
)

func main() {

	var intarr [6]int = [6]int{3, 5, 6, 9, 1, 4}

	slice := intarr[1:3]
	fmt.Println("intarr:", intarr)
	fmt.Println("Slice:", slice)
	fmt.Println("Slice元素个数：", len(slice))
	fmt.Println("Slicen的容量：", cap(slice))

}
