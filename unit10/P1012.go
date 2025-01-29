package main

import (
	"fmt"
	"model"
)

func main() {

	//var s model.Student = model.Student{"lili", 10}
	// s := model.Student{"lili", 10}
	// fmt.Println(s)

	s := model.NewStudent("nana", 10)
	fmt.Println(*s)
}
