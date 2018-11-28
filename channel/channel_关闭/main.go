package main

import (
	"fmt"
)

func main() {

	ch := make(chan int) //创建一个无缓存的Channel

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//不需要写数据时关闭channel
		close(ch)
	}()

	for num := range ch {
		fmt.Println("num=", num)
	}
}

func main01() {

	ch := make(chan int) //创建一个无缓存的Channel

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//不需要写数据时关闭channel
		close(ch)
	}()

	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("num=", num)
		} else {
			fmt.Println("channel关闭")
			break
		}
	}
}
