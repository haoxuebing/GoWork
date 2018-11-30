package main

import "fmt"

func main() {
	var i string = "Hello"
	fmt.Printf("&i=%p\n", &i)
	j := &i
	fmt.Printf("*j=%v,j=%p,&j=%p\n", *j, j, &j)

	k := &j
	fmt.Printf("**k=%v,*k=%v,k=%p,&k=%p\n", *(*k), *k, k, &k)

	var l string = "Hello"
	fmt.Printf("&i=%p,&l=%p\n", &i, &l) //i,l的值虽然相同，但是地址不同
}
