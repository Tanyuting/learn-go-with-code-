package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var ma Teacher
	fmt.Println(ma)
	ma.Name = "mashibing"
	ma.Age = 44
	ma.School = "qingha daxue"
	fmt.Println(ma)
	fmt.Println(ma.Age + 10)
}
