package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("E:/goproject/test01.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件失败", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//writer.WriteString("")
	for i := 0; i < 10; i++ {

		writer.WriteString("你好，\r\n")
	}
	writer.Flush()
	s := os.FileMode(0666).String()
	fmt.Println(s)

}
