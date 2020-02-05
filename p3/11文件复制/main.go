package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFile ...
func CopyFile(srcFileName, dstFileName string) {
	// 判断是否同名
	if srcFileName == dstFileName {
		fmt.Println("目标文件名不能和源文件名相同")
		return
	}
	// 打开源文件
	f, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件
	f2, err3 := os.Create(dstFileName)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	// 完成后自动关闭
	defer f.Close()
	defer f2.Close()
	// 创建缓冲
	b := make([]byte, 1024*6)
	for {

		// 读取源文件
		_, err2 := f.Read(b)
		if err2 != nil {
			// 判断是否到行尾
			if err2 == io.EOF {
				fmt.Println("复制完毕")
				return
			}
			fmt.Println(err2)
			return

		}
		_, err4 := f2.Write(b)
		if err4 != nil {
			fmt.Println(err4)
			return
		}

	}
}

func main() {
	list := os.Args
	if len(list) < 3 {
		fmt.Println("格式：xxx srcFile dstFile")
		return
	}
	sF := list[1]
	dF := list[2]
	CopyFile(sF, dF)
}
