package main

import (
	"fmt"
	"net"
)

func main() {
	listenner, err := net.Listen("tcp", ":8000")

	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}

	defer listenner.Close()

	//阻塞等待用户的连接
	conn, e1 := listenner.Accept()

	if e1 != nil {
		fmt.Println("Accept err:", e1)
		return
	}

	defer conn.Close()

	//接收客户端的数据
	buf := make([]byte, 2048)
	n, e2 := conn.Read(buf)

	if n == 0 {
		fmt.Println("Read err=", e2)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))

}
