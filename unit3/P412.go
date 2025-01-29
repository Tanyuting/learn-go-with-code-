package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue
		}
		fmt.Println(i)
	}
}


package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {

			if i == 2 && j == 2 {
				continue
			}
			fmt.Printf("i: %v, j: %v \n", i, j)
		}
	}
	fmt.Println("OK")
}
