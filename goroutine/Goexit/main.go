package main

import (
	"fmt"
	"runtime"
)

func Test() {
	defer fmt.Println("cccccc")

	runtime.Goexit()

	fmt.Println("dddddddd")
}

func main() {

	go func() {
		fmt.Println("aaaaaaa")

		Test()

		fmt.Println("bbbbbbb")
	}()

	for {

	}

}
