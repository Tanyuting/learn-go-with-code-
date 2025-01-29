package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("E:/goproject/test.txt")

	if err != nil {
		fmt.Println("读取有误", err)

	}

	fmt.Printf("%v", string(content))
}
