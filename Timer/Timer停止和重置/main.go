package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	ok := timer.Reset(1 * time.Second) //重置时间设置为1s

	fmt.Println("ok=", ok)
	<-timer.C
	fmt.Println("时间到")

}

func main01() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，因为定时器的时间到了")
	}()

	timer.Stop() //停止定时器

	for {

	}

}
