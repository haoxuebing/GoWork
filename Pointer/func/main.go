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

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	i, j := 10, 20
	fmt.Printf("交换前i:%d,j:%d\n", i, j)
	swap(&i, &j)
	fmt.Printf("交换后i:%d,j:%d", i, j)
}
