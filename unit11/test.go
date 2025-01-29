package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// 指定要读取的文件路径
	filePath := "E:/goproject/test.txt"

	// 使用ioutil.ReadFile函数读取文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 将文件内容转换为字符串
	fileContent := string(data)

	// 打印文件内容
	fmt.Println(fileContent)
}
