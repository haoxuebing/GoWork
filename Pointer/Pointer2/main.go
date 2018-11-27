package main

import "fmt"

func main() {
	a := 3
	b := 4
	p1 := &a //获取变量a的内存地址，并将其赋值给变量p1
	p2 := &b //获取变量b的内存地址，并将其赋值给变量p2
	fmt.Printf("a的值为 %v, a的指针是 %v ，p1指向的变量的值为 %v\n", a, p1, *p1)
	fmt.Printf("b的值为 %v, b的指针是 %v ，p2指向的变量的值为 %v\n", b, p2, *p2)
	fmt.Println(demo(p1, p2))
	fmt.Printf("a的值为 %v, a的指针是 %v ，p1指向的变量的值为 %v\n", a, p1, *p1)
	fmt.Printf("b的值为 %v, b的指针是 %v ，p2指向的变量的值为 %v\n", b, p2, *p2)

	a = 7
	b = 8
	fmt.Println()
	fmt.Printf("a的值为 %v, a的指针是 %v ，p1指向的变量的值为 %v\n", a, p1, *p1)
	fmt.Printf("b的值为 %v, b的指针是 %v ，p2指向的变量的值为 %v\n", b, p2, *p2)
}

func demo(a, b *int) int {
	*a = 5
	*b = 6
	return *a * *b //这里出现连续的两个*，Go编译器会根据上下文自动识别乘法与两个引用
}
