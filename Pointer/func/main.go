/*
 * @Author: haoxuebing
 * @Date: 2018-08-17 15:46:07
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-08-17 15:47:18
 * @Description: 指针函数，交换两个数的值
 */
package main

import (
	"fmt"
)

//指针是值传递，传入的参数是地址变量的拷贝，所以只改变了变量的值，而主函数中变量的地址不变
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	i, j := 10, 20
	x, y := &i, &j
	fmt.Printf("&x=%p,&y=%p \n", &x, &y)
	fmt.Printf("&x=%p,&y=%p \n", x, y)
	fmt.Printf("&i=%p,&j=%p \n", &i, &j)
	fmt.Printf("交换前i:%d,j:%d\n", i, j)
	swap(&i, &j)
	*x, *y = *y, *x
	fmt.Printf("&i=%p,&j=%p \n", &i, &j)
	fmt.Printf("交换后i:%d,j:%d", i, j)
}
