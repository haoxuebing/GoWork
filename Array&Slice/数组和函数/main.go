/*
 * @Author: haoxuebing
 * @Date: 2018-08-17 16:05:38
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-08-17 16:13:23
 * @Description:函数中传入的数组只是一份拷贝，改变数组的值，原数组不变
 */
package main

import (
	"fmt"
)

func ChangeArr(arr [5]int) {
	arr[0] = 6
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	PrintArr(a)
	ChangeArr(a)
	PrintArr(a)
}

func PrintArr(arr [5]int) {

	fmt.Println("\n打印数组:")
	for _, v := range arr {
		fmt.Print(v, ",")
	}
}
