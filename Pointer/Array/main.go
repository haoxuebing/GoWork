/*
 * @Author: haoxuebing
 * @Date: 2018-08-17 15:47:23
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-08-17 15:54:27
 * @Description:指针数组
 */

package main

import (
	"fmt"
)

func main() {

	a := []int{10, 20, 30}

	var ptr [3]*int

	for i := 0; i < 3; i++ {
		ptr[i] = &a[i]
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d]=%d\n", i, *ptr[i])
	}

}
