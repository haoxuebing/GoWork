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
	arr[1] = 20
	arr = append(arr, 8) //此处的元素并没有增加到切片
}

func ChangeArr2(arr *[]int) {
	(*arr)[0] = 6          //因为传入的是切片指针，所以所有使用切片的地方都必须是指针类型的切片（这里与数组不同）
	*arr = append(*arr, 7) //此处的元素添加进了切片
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	PrintArr(a)
	ChangeArr(a)
	PrintArr(a)
	ChangeArr2(&a)
	PrintArr(a)
}

func PrintArr(arr []int) {

	fmt.Println("\n打印切片:")
	for _, v := range arr {
		fmt.Print(v, ",")
	}
}
