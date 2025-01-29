package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := strings.Replace("goandjavagogo", "go", "golnag", -1)
	fmt.Println(str1)

	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)

}
