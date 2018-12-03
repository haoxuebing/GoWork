package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("resp.Status=", resp.Status)

	buf := make([]byte, 1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)

		if n == 0 {
			fmt.Println("read err=", err)
			break
		}
		tmp += string(buf[:n])
	}

	fmt.Println("tmp=", tmp)

}
