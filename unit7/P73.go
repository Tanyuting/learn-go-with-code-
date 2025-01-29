package main

import "fmt"

func main() {
	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请输入学生%d的成绩", i+1)

		fmt.Scanln(&scores[i])
	}
	for i := 0; i < len(scores); i++ {

		fmt.Printf("弟%d个学生的成绩:  %d\n", i+1, scores[i])

	}
	fmt.Printf("11111111111111111")
	for key, value := range scores {
		fmt.Printf("第%d个学生的成绩: %d\n", key+1, value)
	}
}
