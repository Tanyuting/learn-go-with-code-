package main

import (
	"fmt"
	"strings"
)

func main() {
	var a map[int]string

	a = make(map[int]string, 10)
	a[1123123] = "ddd"
	a[1233455] = "www"

	fmt.Println(a)

	b := make(map[int]string)
	b[12324] = "张三"
	b[23344] = "李四"
fmt.Println(b)
	c := map[int]string{

		12345:"zhangsn"
		23445:"lsi sidd"


	}
	c[2342324] = "wangwu"
	fmt.Println(c)
}
