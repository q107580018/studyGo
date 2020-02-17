package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

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

// SpyderOneJoy 爬取单个笑话
func SpyderOneJoy(url string) (title, text string, err error) {
	res, err := HTTPGet(url)
	if err != nil {
		fmt.Println("HTTPGet err:", err)
		return
	}
	re := regexp.MustCompile(`<h1 class="entry-title">(?s:(.*?))</h1>`)
	title = re.FindStringSubmatch(string(res))[1]
	title = strings.TrimSpace(title)
	title = strings.Trim(title, "【搞笑段子】")
	title = strings.Trim(title, " - 嗡啪网")
	re = regexp.MustCompile(`<meta name="description" content="(?s:(.*?))"/>`)
	texts := re.FindStringSubmatch(string(res))
	if len(texts) != 2 {
		return
	}
	text = texts[1]
	text = strings.Replace(text, "<br />", "", -1)
	return
}

// SpiderPage 爬取网页
func SpiderPage(i int) {
	JoyMap := make(map[string]string)
	fmt.Println("正在爬取第", i, "页网页...")

	url := fmt.Sprintf("https://wengpa.com/duanzi/page/%d/", i)
	res, err := HTTPGet(url)
	if err != nil {
		fmt.Println("HTTPGet err:", err)
		return
	}
	// 筛选出段子链接
	re := regexp.MustCompile(`<a href="(.*?)" rel="bookmark" target="_blank" style="z-index:-100" >`)

	result := re.FindAllSubmatch(res, -1)

	for _, data := range result {
		// fmt.Println(string(data[1]))
		url := string(data[1])
		title, text, err := SpyderOneJoy(url)
		if err != nil {
			fmt.Println("spyderOneJoy err:", err)
		}
		if title != "" && text != "" {
			JoyMap[title] = text
		}
	}
	WriteToFile(i, JoyMap)

}

// WriteToFile 保存到本地
func WriteToFile(i int, JoyMap map[string]string) {
	f, err := os.Create(fmt.Sprintf("%d.txt", i))
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()
	for title, text := range JoyMap {
		f.WriteString(title + "\n")
		f.WriteString(text + "\n\n")
	}
}

// DoWork 爬取网页
func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)
	for i := start; i <= end; i++ {
		SpiderPage(i)
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
