package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var page = make(chan int)

// DownloadToFile 将爬取的网页保存
func DownloadToFile(res []byte, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()
	f.Write(res)
}

// HTTPGet 获取网页信息
func HTTPGet(url string) (res []byte, err error) {
	res = make([]byte, 0, 4096)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 && err == io.EOF {
			break
		}
		res = append(res, buf[:n]...)
	}
	return
}

// SpiderPage 爬取网页
func SpiderPage(i int) {
	fmt.Println("正在爬取第", i, "页网页...")
	pn := strconv.Itoa(50 * (i - 1))
	url := "https://tieba.baidu.com/f?kw=golang&ie=utf-8&pn=" + pn
	res, err := HTTPGet(url)
	if err != nil {
		fmt.Println("HTTPGet err:", err)
		return
	}
	fileName := strconv.Itoa(i) + ".html"
	DownloadToFile(res, fileName)
	page <- i
}

// DoWork 爬取网页
func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)
	for i := start; i <= end; i++ {
		go SpiderPage(i)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页网页爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Print("请输入起始页：")
	fmt.Scan(&start)
	fmt.Print("请输入终止页：")
	fmt.Scan(&end)
	if end < start {
		fmt.Println("终止页不能小于起始页")
		return
	}
	DoWork(start, end)

}
