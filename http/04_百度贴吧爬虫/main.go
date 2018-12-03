package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var start, end int
	fmt.Println("请输入请求启始页:")
	fmt.Scan(&start)
	fmt.Println("请输入请求终止页:")
	fmt.Scan(&end)

	DoWork(start, end)
}

func DoWork(start, end int) {
	fmt.Printf("正在爬去 %d 到 %d 的页面\n", start, end)

	page := make(chan int)

	for i := start; i < end; i++ {
		go SpiderPage(i, page)
	}

	for i := start; i < end; i++ {
		fmt.Printf("第%d个页面爬去完成\n", <-page)
	}
}

func SpiderPage(i int, page chan int) {
	url := "http://tieba.baidu.com/f?kw=石家庄学院&ie=utf-8&pn=" + strconv.Itoa(50*i)
	fmt.Println("url=", url)

	//爬取网页内容
	rlt, err := HttpGET(url)

	if err != nil {
		fmt.Println("HttpGet err =", err)
		return
	}

	//把内容写入到文件

	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)

	if err1 != nil {
		fmt.Println("os.Create err1=", err1)
		return
	}

	f.WriteString(rlt)
	f.Close()

	page <- i
}

//爬去网页内容
func HttpGET(url string) (result string, err error) {

	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)

	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			// fmt.Println("resp.Body.Read err=", err)
			break
		}

		result += string(buf[:n])
	}

	return result, err
}
