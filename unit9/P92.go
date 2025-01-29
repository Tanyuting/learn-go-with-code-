package main

import (
	"fmt"
)

func main() {

	b := make(map[int]string)
	b[12324] = "张三"
	b[23344] = "李四"

	b[12324] = "wangbin"

	fmt.Println(b)

	delete(b, 23344)
	fmt.Println(b)

	value, flag := b[12324]
	fmt.Println(value)
	fmt.Println(flag)
}
