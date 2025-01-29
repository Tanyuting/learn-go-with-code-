package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	//var t Teacher = Teacher{"tanyuting", 333, "dddd"}
	//fmt.Println(t)
	var t *Teacher = new(Teacher)
	(*t).Name = "xxx"
	(*t).Age = 3333
	(*t).School = "xxxxx"
	t.School = "dddd"
	fmt.Println(*t)

}
