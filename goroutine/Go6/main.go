package main

import (
	"fmt"
	"strconv"
	"time"
)

func NewTask(str string) {

	for {
		fmt.Println("打印:", str)
		time.Sleep(time.Second)
	}
}

func main() {
	for i := 0; i <= 1000; i++ {
		go NewTask(strconv.Itoa(i))
	}

	NewTask("AAAAAA")

}
