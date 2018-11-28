package main

import (
	"fmt"
	"strconv"
)

func MyPrintf(args ...interface{}) {
	for _, arg := range args {

		fmt.Println(fmt.Sprintf("%T", arg))
		fmt.Println(strconv.Itoa(123))

		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value")
		case string:
			fmt.Println(arg, "is a string value")
		case int64:
			fmt.Println(arg, "is an  int64 value")
		default:
			fmt.Println(arg, "is an unknown type")
		}

	}

}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "Hello"
	var v4 float32 = 1.33
	MyPrintf(v1, v2, v3, v4)

}
