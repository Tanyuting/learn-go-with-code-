package main

import (
	"fmt"
)

func main() {

	//var intChan1 chan int
	var intChan2 chan<- int //只写
	intChan2 = make(chan int, 3)
	intChan2 <- 20
	//num := <-intChan2
	fmt.Println("intchan2", intChan2)

	var intChan3 <-chan int //只读
	if intChan3 != nil {
		num1 := <-intChan3
		fmt.Println("num1", num1)
	}
	//intChan3 <- 30
}
