package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "golang你好"
	fmt.Println((len(str)))

	for i, value := range str {
		fmt.Printf("索引为: %d, 具体值：%c \n", i, value)
	}
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c, \n", r[i])
	}

	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)

	str1 := strconv.Itoa(88)
	fmt.Println(str1)
	count := strings.Count("golangandjave", "ga")
	fmt.Println(count)

	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)
	fmt.Println("hello" == "HELLO")

	fmt.Println(strings.Index("golangandjave", "ga"))
}
