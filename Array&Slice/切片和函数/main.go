/*
 * @Author: haoxuebing
 * @Date: 2018-08-17 16:13:54
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-08-17 16:16:02
 * @Description: 切片是引用类型，在函数中改变切片的值，原切片改变
 */

package main

import (
	"fmt"
)

func ChangeArr(arr []int) {
	arr[0] = 6
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	PrintArr(a)
	ChangeArr(a)
	PrintArr(a)
}

func PrintArr(arr []int) {

	fmt.Println("\n打印数组:")
	for _, v := range arr {
		fmt.Print(v, ",")
	}
}
