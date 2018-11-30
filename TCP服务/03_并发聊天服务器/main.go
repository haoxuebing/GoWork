package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用户发数据的管道
	Name string
	Addr string
}

//保存在线用户
var onlineMap map[string]Client

var messgae = make(chan string)

func main() {
	listenner, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err= ", err)
		return
	}
	defer listenner.Close()

	//新开一个协程，转发消息，只要有消息来，变量map给每个成员都发送此消息
	go Manager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("listenner.Accept err =", err)
			continue
		}

		go HandleConn(conn) //处理用户连接

	}
}

func HandleConn(conn net.Conn) {

	defer conn.Close()
	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()

	cli := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = cli

	//新开一个协发消息
	go WriteMsgToClient(cli, conn)

	//广播某个在线
	// messgae <- "[" + cli.Addr + "]" + cli.Name + ": login"
	messgae <- MakeMsg(cli, "login")

	//提示我是谁
	cli.C <- MakeMsg(cli, "I am here")

	isQuit := make(chan bool)  //用户主动退出
	hasData := make(chan bool) //是否有数据发送

	//新开一个协程接收用户发过来的数据
	go func() {
		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)
			if n == 0 { //对方断开，或者出问题
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:n-1])

			if msg == "who" {
				//遍历Map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}

			} else if len(msg) > 7 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
				//转发此内容
				messgae <- MakeMsg(cli, string(msg))
			}
			hasData <- true
		}
	}()

	for {
		//通过elect 检测channel的流动

		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)
			messgae <- MakeMsg(cli, "login out")
			return
		case <-hasData:

		case <-time.After(60 * time.Second): //60s后超时处理
			delete(onlineMap, cliAddr)
			messgae <- MakeMsg(cli, "time out leave out")
		}
	}
}

func MakeMsg(cli Client, msg string) string {
	buf := "[" + cli.Addr + "]" + cli.Name + ": login"
	return buf
}

func Manager() {

	onlineMap = make(map[string]Client)

	for {
		msg := <-messgae

		//给每个成员都发送消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {

	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}

}
