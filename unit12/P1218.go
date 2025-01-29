package main

import (
	"fmt"
	"time"
)

func main() {

	intchan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 2)
		intchan <- 20
	}()

	stringchan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 5)
		stringchan <- "testtttt"
	}()
	//fmt.Println(<-intchan)

	select {

	case v := <-intchan:
		fmt.Println("intchan:", v)
	case v := <-stringchan:
		fmt.Println("stringchan:", v)

	default:
		fmt.Println("防止select被阻塞")
	}
}
