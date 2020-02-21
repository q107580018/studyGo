package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CopyFile ...
func CopyFile(srcFileName, dstFileName string) (written int64, err error) {
	// 判断是否同名
	if srcFileName == dstFileName {
		fmt.Println("目标文件名不能和源文件名相同")
		return
	}
	// 打开源文件
	f, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	// 创建新文件
	f2, err := os.Create(dstFileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	// 完成后自动关闭
	defer f.Close()
	defer f2.Close()
	src := bufio.NewReader(f)
	dst := bufio.NewWriter(f2)
	defer dst.Flush() // 重要！
	return io.Copy(dst, src)

}

func main() {
	list := os.Args
	if len(list) < 3 {
		fmt.Println("格式：xxx srcFile dstFile")
		return
	}
	sF := list[1]
	dF := list[2]
	_, err := CopyFile(sF, dF)
	if err == nil {
		fmt.Println("拷贝完成")
	}
}
