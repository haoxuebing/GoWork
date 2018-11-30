/*
 * @Author: haoxuebing
 * @Date: 2018-11-30 15:15:08
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-11-30 15:15:55
 * @Description: TCP通信 服务端
 */
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户链接
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("err=", err)
			return
		}

		//接收用户请求
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	//函数调用完毕，自动关闭
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Printf("addr: %s connect successful! \n", addr)

	buf := make([]byte, 2048)

	//读取用户数据
	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("err=", err)
			return
		}

		msg := string(buf[:n])
		fmt.Printf("[%s]: %s\n", addr, msg)

		if strings.Contains(msg, "exit") {
			fmt.Println(addr, " exit")
			return
		}

		//把数据转化为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(msg)))
	}
}
