package main

import "fmt"

func main() {
	var str string = "hello golang"
	//for i := 0; i < len(str); i++ {
	//fmt.Printf("%c \n", str[i])

	for i, value := range str {
		fmt.Printf("索引：%d, 具体值：%c \n", i, value)
	}
}
