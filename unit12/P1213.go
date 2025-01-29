package main

import "fmt"

func main() {

	var intchan chan int
	intchan = make(chan int, 3)

	intchan <- 10
	intchan <- 20

	close(intchan)

	num := <-intchan
	fmt.Println(num)
}
