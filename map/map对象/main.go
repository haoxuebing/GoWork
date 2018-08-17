/*
 * @Author: haoxuebing
 * @Date: 2018-08-17 18:17:26
 * @Last Modified by: haoxuebing
 * @Last Modified time: 2018-08-17 18:20:04
 * @Description: 主要用interface实现泛型的效果

 switch v.(type){
 case int:
	//to do
 case string:
	//to do
 }


obj.(map[string]string)  //把inteface类型的obj 转化为 map[string]string

*/
package main

import "fmt"

type PersonInfo struct {
	ID   string
	Name string
}

func main() {
	//var PersonDB map[string]PersonInfo
	//PersonDB = make(map[string]PersonInfo)

	myMap := make(map[string]PersonInfo)

	myMap["001"] = PersonInfo{
		ID:   "001",
		Name: "张三",
	}

	PrintMap(myMap)

	myDic := make(map[string]string)

	myDic["aa"] = "asdf"
	myDic["bb"] = "qwer"

	PrintMap(myDic)

	j, k := 1, "hello"
	fmt.Println(j, "\t", k)
}

func PrintMap(m interface{}) {

	switch m.(type) {
	case map[string]string:
		a, _ := m.(map[string]string)
		for i, v := range a {
			fmt.Println(i, "value:", v)
		}
	case map[string]PersonInfo:
		PrintMapPersonInfo(m.(map[string]PersonInfo))
	}

}

func PrintMapPersonInfo(m map[string]PersonInfo) {
	for i, v := range m {
		fmt.Println(i, "Persoinfo:", v)
	}
}
