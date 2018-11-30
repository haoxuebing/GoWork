package main

import (
	"fmt"
)

//面向过程
func Add01(a, b int) int {
	return a + b
}

//面向对象
type long int

func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {

	rlt := Add01(1, 2)
	fmt.Println(rlt)

	var a long = 2
	fmt.Println(a.Add02(3))
	//面向对象只是换了一种表现形式

}
