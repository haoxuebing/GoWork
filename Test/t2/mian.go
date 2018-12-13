package main

import (
	"fmt"
)

func main() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	fmt.Println(arr[1:])
	fmt.Println(arr)

	fmt.Println(arr[:len(arr)])

	fmt.Println(arr)

}
