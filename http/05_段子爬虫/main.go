package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	fmt.Printf("准备爬取第%d到%d的网页\n", start, end)

	page := make(chan int)

	for i := start; i < end; i++ {
		//爬去页面
		go SpiderPage(i, page)
	}

	for i := start; i < end; i++ {
		fmt.Printf("第%d个网页爬完了\n", <-page)
	}

	fmt.Println("所有网页结束")

}

func SpiderPage(i int, page chan int) {
	url := "https://www.pengfu.com/index_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬去第%d个页面,url:%s\n", i, url)

	rlt, err := HttpGet(url)

	if err != nil {
		fmt.Println("HttpGet err=", err)
		return
	}

	//正则表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)

	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	joyUrls := re.FindAllStringSubmatch(rlt, -1)

	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	//取网址
	for _, data := range joyUrls {

		title, content, err := SpiderOneJoy(data[1])
		if err != err {
			fmt.Println("SpiderOneJoy err=", err)
			continue
		}
		// fmt.Println("title=", title)
		// fmt.Println("content=", content)
		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)

	}

	//把内容写入到文件
	StoreJoyToFile(i, fileTitle, fileContent)
	page <- i
}

func StoreJoyToFile(i int, fileTitle, fileContent []string) {
	//新建文件
	f, _ := os.Create(strconv.Itoa(i) + ".txt")

	defer f.Close()
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		//写标题
		f.WriteString(fileTitle[i] + "\n")
		//写内容
		f.WriteString(fileContent[i] + "\n")
		f.WriteString("============================\n")
	}

}

func SpiderOneJoy(url string) (title, content string, err error) {
	rlt, err := HttpGet(url)

	//取关键信息
	re := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	tmpTitle := re.FindAllStringSubmatch(rlt, 1)
	for _, data := range tmpTitle {
		title = strings.Replace(data[1], "\n", "", -1)
		title = strings.Replace(title, " ", "", -1)
		break
	}

	//取内容
	re = regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	tmpContent := re.FindAllStringSubmatch(rlt, 1)
	for _, data := range tmpContent {
		content = strings.Replace(data[1], "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, " ", "", -1)

		break
	}
	return

}

func HttpGet(url string) (rlt string, err error) {

	resp, err := http.Get(url)

	buf := make([]byte, 1024)

	for {

		n, _ := resp.Body.Read(buf)

		if n == 0 {
			break
		}

		rlt += string(buf[:n])

	}

	return rlt, err
}
