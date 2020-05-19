package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("/Users/qiuwen/Nutstore Files/我的坚果云/go_src/q107580018/studyGo/p3/18fileInfo/aaa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	// 获取文件名
	fmt.Println("文件名:", fileInfo.Name())
	// 获取文件大小
	fmt.Println("文件大小:", fileInfo.Size())
	// 是否为目录
	fmt.Println(fileInfo.IsDir())
	// 修改时间
	fmt.Println(fileInfo.ModTime())
	// 权限
	fmt.Println(fileInfo.Mode()) // -rwx------
}
