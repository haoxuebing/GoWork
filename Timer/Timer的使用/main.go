package main

import (
	"fmt"
	"time"
)

func main() {
	<-time.After(2 * time.Second)
	fmt.Println("时间到")
}

func main03() {
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
}

func main02() {
	timer := time.NewTimer(1 * time.Second)

	<-timer.C
	fmt.Println("时间到")

}

func main01() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2秒后，往timer.C写数据，有数据后就可以读取
	t := <-timer.C //channel 没有数据前后阻塞
	fmt.Println("t=", t)

}
