package main

import "fmt"

func main() {

	var intchan chan int
	intchan = make(chan int, 3)

	fmt.Printf("intchan的值：%v", intchan)

	intchan <- 10
	num := 20
	intchan <- num
	intchan <- 40

	num1 := <-intchan
	num2 := <-intchan
	num3 := <-intchan

	fmt.Println(num1, num2, num3)

	fmt.Printf("管道的实际长度：%v，管道的容量是：%v", len(intchan), cap(intchan))
}
