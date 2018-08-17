package main

import (
	"fmt"
	"reflect"
)

func main() {

	//声明数组
	var Arr [5]int

	for i := 0; i < len(Arr); i++ {
		Arr[i] = i * 5
	}
	fmt.Println("打印声明数组：")
	PrintArr(Arr[:])

	//数组
	myArr := []int{1, 2, 3, 4}
	fmt.Println("打印myArr:")
	PrintArr(myArr[:])
	//数组创建切片
	mySlice1 := myArr[:]
	fmt.Println("打印mySlice1:")
	PrintArr(mySlice1)

	//make创建切片 指定接片大小 避免内存重复申请
	mySlice2 := make([]int, 10)
	mySlice2 = append(mySlice2, 2, 3)
	fmt.Println("打印mySlice2:")
	PrintArr(mySlice2)

	//声明空切片
	var nSlice []int
	fmt.Println("打印nSlice:s")
	nSlice = append(nSlice, 9, 8, 7, 6)
	PrintArr(nSlice)

	strSlice := []string{"abc", "def"}
	fmt.Println("打印string数组：")
	PrintStringArr(strSlice)
	fmt.Println("改变string数组：")
	//PrintStringArr(ChangeSlice(strSlice))
	PrintEveryArr(strSlice)

	MyTest()

}

func PrintArr(arr []int) {
	for i, v := range arr {
		fmt.Println(i, "\t", v)
	}
}

func PrintStringArr(arr []string) {
	for i, v := range arr {
		fmt.Println(i, "\t", v)
	}
}

func PrintEveryArr(arr interface{}) {
	//for i, v := range arr {
	//	fmt.Println(i, "\t", v)
	//}
	fmt.Println(arr)
	//j,k:=arr.(type)
	//fmt.Println("arr 类型为：", reflect.TypeOf(arr))

	str := reflect.TypeOf(arr)

	//type myType str
	fmt.Println("arr 类型为：", str)

	var newArr []string = arr.([]string)
	for i, v := range newArr {
		fmt.Println(i, "\t", v)
	}
}

func ChangeSlice(s []string) []string {
	s[0] = "qwe"
	s[1] = "asd"
	s = append(s, "zxc")
	return s
}
