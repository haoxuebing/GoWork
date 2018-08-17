/*
 * @Author: haoxuebing
 * @Date: 2018-08-17 17:28:45
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-08-17 17:29:33
 * @Description: 函数中改变传入的Map,原Map值会改变
 */
package main

import (
	"fmt"
)

func ChangeMap(m map[string]string) {

	m["张"] = "张三"

	m["李"] = "李四"

}

func main() {
	m := make(map[string]string)
	ChangeMap(m)
	PrintMap(m)
}

func PrintMap(m map[string]string) {

	for k, v := range m {
		fmt.Printf("k:%s,v:%s\n", k, v)
	}
}
