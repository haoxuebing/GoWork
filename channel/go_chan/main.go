package main

import (
	"fmt"
)

func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 2
}

func main() {
	chs := make([]chan int, 20)
	for i := 0; i < 20; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
}
