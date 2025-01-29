package main

import (
	"fmt"
)

func main() {

	b := make(map[int]string)
	b[12324] = "张三"
	b[23344] = "李四"

	b[123654] = "wangbin"
	fmt.Println(len(b))
	for k, v := range b {
		fmt.Printf("keyis ：%v value is%v \t", k, v)
	}

	a := make(map[string]map[int]string)

	a["class1"] = make(map[int]string, 3)
	a["class1"][220095566] = "lulu"
	a["class1"][220033550] = "lili"
	a["class1"][200334444] = "xxxx"

	a["class2"] = make(map[int]string, 3)
	a["class2"][220095566] = "lulu"
	a["class2"][220033550] = "lili"
	a["class2"][200334444] = "xxxx"

	for k1, v1 := range a {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("学号2：%v，姓名：%v \t", k2, v2)
		}
		fmt.Println()
	}
}
