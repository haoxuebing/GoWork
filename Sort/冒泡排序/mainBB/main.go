package main

import (
	"../SortFun"
	"fmt"
)

func main() {
	arr := []int{2, 3, 7, 9, 0}
	// testBB.BubbleSort(arr)
	SortFun.QuckSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
