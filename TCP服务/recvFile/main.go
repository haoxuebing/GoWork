package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	defer listenner.Close()

	//阻塞等待用户连接
	for {
		conn, e2 := listenner.Accept()

		if e2 != nil {
			fmt.Println("Listenner.Accept err =", e2)
			return
		}

		buf := make([]byte, 1024)
		n, e3 := conn.Read(buf)
		if e3 != nil {
			fmt.Println("conn.Read err =", err)
			return
		}

		fileName := string(buf[:n])

		conn.Write([]byte("ok"))

		//接收文件内容
		RecvFile(fileName, conn)
	}

}

func RecvFile(fileName string, conn net.Conn) {
	f, _ := os.Create(fileName)

	buf := make([]byte, 1024*4)

	//接收多少写多少
	for {
		n, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
				return
			} else {
				fmt.Println("conn.Read err=", err)
			}
			return
		}

		if n == 0 {
			fmt.Println("文件接收完毕n=0")
			break
		}

		f.Write(buf[:n])
	}

}
