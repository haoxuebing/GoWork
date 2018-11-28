package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func PrinteT(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

//pserson1阻塞完之后 pserson2执行
func Person1() {
	PrinteT("Hello")
	ch <- 666
}

func Person2() {
	<-ch
	PrinteT("World")
}

func main() {
	go Person1()
	go Person2()

	for {

	}
}
