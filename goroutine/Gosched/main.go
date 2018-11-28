package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片，先让别的协程执行，别的协程执行完，在回来执行此协程
		runtime.Gosched()
		fmt.Println("Hello")
	}
}