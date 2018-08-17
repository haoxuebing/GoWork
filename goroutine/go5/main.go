package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
}

func main() {
	go ready("js", 2)
	go ready("C#", 1)
	fmt.Println("I'm waiting")
	time.Sleep(3 * time.Second)
}
