package main

import "fmt"

func main() {

	var age int

	var name string

	var score float32

	var isVIP bool

	fmt.Println("enter age name score VIP ")
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)
	fmt.Printf("stdunt age: %v, name: %v, score: %v, vip %v", age, name, score, isVIP)
}
