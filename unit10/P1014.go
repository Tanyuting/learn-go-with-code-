package main

import (
	"fmt"
	"model"
)

func main() {

	p := model.Newperson("lili")
	p.Setage(20)

	fmt.Println(p.Name)
	fmt.Println(p.Getage())
	fmt.Println(*p)

}
