package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WriteFile ...
func WriteFile(path string) {
	f, err := os.Create(path)
	defer f.Close() // 使用完毕，自动关闭文件
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		var str string
		str = fmt.Sprintf("i = %d\n", i)
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// ReadFile ...
func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("error =", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*2) // 2k大小
	_, err = f.Read(buf)
	if err == io.EOF { // 文件结尾
		fmt.Println("文件读取结束")
	}
	fmt.Println(string(buf))
}

// ReadFileLine 按行读取
func ReadFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(line[:len(line)-1])
	}
}

func main() {
	path := "./demo.txt"
	// WriteFile(path)
	// ReadFile(path)
	ReadFileLine(path)
}
