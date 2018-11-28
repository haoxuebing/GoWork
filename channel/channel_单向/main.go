package main

import (
	"fmt"
)

func producter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num=", num)
	}
}

func main() {
	ch := make(chan int)

	go producter(ch)

	consumer(ch)
}

func main01() {

	ch := make(chan int)

	//双向channel隐式转化为单向channel
	var wch chan<- int = ch //只能写，不能读
	wch <- 666              // 写

	var rch <-chan int = ch //只能读，不能写
	<-rch                   //读

}
