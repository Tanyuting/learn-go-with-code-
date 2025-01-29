package main

import "fmt"

func main() {
	var intchan chan int
	intchan = make(chan int, 100)

	for i := 0; i < 100; i++ {

		intchan <- i
	}
	close(intchan)
	for v := range intchan {
		fmt.Println("value =", v)
	}
}
