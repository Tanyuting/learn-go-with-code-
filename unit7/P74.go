package main

import "fmt"

func main() {
	var aar1 [3]int = [3]int{3,6,9}
	fmt.Println(aar1)

	var aar2 =[3]int{1,4,7}
	fmt.Println(aar2)

	var aar3 = [...]{4,5,6,7}
	fmt.Println(aar3)

	var aar4 = [...]int{2:66,0:33,1:99,3:88}
	fmt.Println(aar4)
}