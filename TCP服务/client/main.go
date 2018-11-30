/*
 * @Author: haoxuebing
 * @Date: 2018-11-30 15:15:26
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-11-30 15:26:09
 * @Description: TCP通信 客户端
 */
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer conn.Close()

	//从键盘输入内容，给服务器发送内容
	go func() {
		str := make([]byte, 2048)
		for {
			n, err := os.Stdin.Read(str)

			if err != nil {
				fmt.Println("os.Stdin.err=", err)
				return
			}
			//把输入的内容发送给服务端
			ss := strings.Replace(string(str[:n]), "\n", "", -1)
			conn.Write([]byte(ss))
		}
	}()

	//接收服务器回复的数据
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf) //接收服务器的请求
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println(string(buf[:n])) //打印接收到的内容
	}

}
