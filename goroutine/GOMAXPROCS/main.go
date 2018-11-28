package main

import (
	"fmt"
	// "runtime"
)

func main() {
	// n := runtime.GOMAXPROCS(2) //设置2个核执行，不设置默认为系统核数

	// fmt.Println(n)

	for {
		go fmt.Print(1)

		fmt.Print(0)
	}
}
