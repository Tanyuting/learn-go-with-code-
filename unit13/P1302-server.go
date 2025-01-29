package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}
func main() {
	fmt.Println("服务器启动了。。。")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {

		fmt.Println("监听failed，err:", err)
		return
	}

	for {
		conn, err2 := listen.Accept()
		if err2 != nil {

			fmt.Println("客户端的等待failed，err2", err2)

		} else {
			fmt.Printf("等待连接成功， con=%v, 收到的客户端信息： %v \n", conn, conn.RemoteAddr().String())
		}
		//准备一个协程
		go process(conn)
	}
}
