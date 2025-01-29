package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 4; i++ {

		fmt.Println("hello glang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go test()

	for i := 1; i <= 10; i++ {

		fmt.Println("hello python" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
