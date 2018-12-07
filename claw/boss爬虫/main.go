package main

import (
	// "encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	// "github.com/garyburd/redigo/redis"
	// "log"
	// "math/rand"
	"net/http"
	// "net/url"
	// "os"
	// "strconv"
	// "strings"
	// "time"
)

func main() {
	getPage()
}

func getPage() {
	boss := "https://www.zhipin.com/job_detail/?query=golang&scity=101010100&industry=&position="
	req, _ := http.NewRequest("GET", boss, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	cli := &http.Client{}
	res, _ := cli.Do(req)

	dom, _ := goquery.NewDocumentFromResponse(res)

	dom.Find(".job-list ul li").Each(func(i int, ctx *goquery.Selection) {

		title := ctx.Find(".job-title").Text()
		company := ctx.Find(".company-text .name a").Text()
		salery := ctx.Find(".info-primary .red").Text()
		dtl, _ := ctx.Find(".info-primary .name a").Attr("href")

		fmt.Println(title + ":" + company + ":" + salery + ": https://www.zhipin.com" + dtl)

	})

	// buf := make([]byte, 1024)
	// tmp := ""
	// for {
	// 	n, _ := res.Body.Read(buf)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	tmp += string(buf[:n])
	// }
	// fmt.Println(tmp)

}
