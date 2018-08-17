package main

import (
	"fmt"
)

func main() {
	// var n []int
	n := make([]int, 1)
	n = append(n, 1)
	n = append(n, 2)

	fmt.Println(n)

	array := [4]int{10, 20, 30, 40}

	slice := array[0:2]

	newSlice := append(slice, 50)

	newSlice[1] += 1

	fmt.Println(newSlice)
}
