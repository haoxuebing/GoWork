package main

import (
	"fmt"
	"time"
)

func PrinteT(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func Person1() {

	PrinteT("Hello")

}

func Person2() {
	PrinteT("World")
}

func main() {

	go Person1()
	go Person2()

	for {

	}
}
