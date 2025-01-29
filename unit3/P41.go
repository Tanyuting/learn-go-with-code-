package main

import "fmt"

func main() {

	var age int
	fmt.Println("enter age:")
	fmt.Scanln(&age)

	var name string
	fmt.Println("enter name:")
	fmt.Scanln(&name)

	var score float32
	fmt.Println("enter score:")
	fmt.Scanln(&score)

	var isVIP bool
	fmt.Println("whether is VI:")
	fmt.Scanln(&isVIP)

	fmt.Printf("student age:%v, name:%v,score: %v, VIP or not: %v", age, name, score, isVIP)

}
