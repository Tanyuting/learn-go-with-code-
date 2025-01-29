package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动。。")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		//连接失败
		fmt.Println("客户端连接失败：err：", err)
		return
	}
	fmt.Println("连接successful", conn)

	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败，err", err)
	}

	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("连接失败，err", err)
	}
	fmt.Printf("终端数据通过客户发送成功，一共发送了 %v 字节数据，并退出", n)
}
