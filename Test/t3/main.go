package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	// dictionary := make(map[int]string)
	dictionary := make(map[int]string, 10000000)
	for i := 0; i < 10000000; i++ {
		dictionary[i] = "hello" + strconv.Itoa(i)
		// dictionary[i] = "hello"
	}
	duration := time.Since(start)
	fmt.Println("time used:", duration)
	time.Sleep(time.Second * 5)
}
