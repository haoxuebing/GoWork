package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件:")
	var path string
	fmt.Scan(&path)

	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.stat err=", err)
		return
	}

	//主动连接服务器

	conn, e2 := net.Dial("tcp", "127.0.0.1:8000")
	if e2 != nil {
		fmt.Println("net.Dial e2 =", e2)
		return
	}

	defer conn.Close()

	//给接收方发送文件名

	_, e3 := conn.Write([]byte(info.Name()))

	if e3 != nil {
		fmt.Println("e3=", e3)
		return
	}

	buf := make([]byte, 1024)
	n, e4 := conn.Read(buf)
	if e4 != nil {
		fmt.Println("e4=", e4)
		return
	}

	if "ok" == string(buf[:n]) {
		SendFile(path, conn)
	}

}

func SendFile(path string, conn net.Conn) {

	//以只读方式打开
	f, err := os.Open(path)
	if err != err {
		fmt.Println("os.Open err=", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)

	//读文件内容，多次发送

	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
				return
			} else {
				fmt.Println("f.Read err =", err)
			}

			return
		}

		//发送内容
		conn.Write(buf[:n])
	}

}
