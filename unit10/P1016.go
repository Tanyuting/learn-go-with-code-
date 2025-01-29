package main

import (
	"fmt"
)

type Animal struct {
	Age    int
	Weight float32
}

func (an *Animal) shout() {

	fmt.Println("我可以大叫")
}

func (an *Animal) ShowInfo() {

	fmt.Printf("动物的年龄： %v, 动物的体重： %v", an.Age, an.Weight)
}

type Cat struct {
	Animal
}

func (c *Cat) scratjc() {

	fmt.Println("小猫挠你")
}

func main() {
	cat := &Cat{}
	cat.Age = 3
	cat.Weight = 10.11
	cat.ShowInfo()
	cat.shout()
	cat.scratjc()
}
